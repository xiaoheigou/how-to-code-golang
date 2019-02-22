package runner

//How to use channel to check the runtime of the programm
import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

type Runner struct {
	interrupt chan os.Signal

	complete chan error

	timeout <-chan time.Time

	tasks []func(int, *sync.WaitGroup)
}

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add : add tasks to Runner
func (r *Runner) Add(tasks ...func(int, *sync.WaitGroup)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}
func (r *Runner) run() error {
	// check got interrupt
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan int)
	inter := make(chan error, 1)
	go func() error {
		for {
			select {
			case <-ticker.C:
				// do stuff
				if r.gotInterrupt() {
					inter <- ErrInterrupt
				}
			case <-quit:
				ticker.Stop()
				return nil
			}
		}
	}()

	// run task
	total := len(r.tasks)
	fmt.Println(total, "**********")
	wg.Add(total)
	fmt.Println(wg, "------")
	go func(*sync.WaitGroup) {
		for id, task := range r.tasks {
			go task(id, &wg)
		}
		wg.Wait()
		fmt.Println("++++++++++complete")
		quit <- 1
	}(&wg)
	// judge complete or interrupt
	fmt.Println("83")
	select {
	case <-quit:
		fmt.Println("quit")
		return nil
	case <-inter:
		return ErrInterrupt
	}
}
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
