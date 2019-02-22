package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/xiaoheigou/how-to-code/07-go-in-action-currency/runner.1/runner"
)

const timeout = 5 * time.Second

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {

	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}
	fmt.Println("complete -_-")
}
func createTask() func(int, *sync.WaitGroup) {
	return func(id int, wg *sync.WaitGroup) {

		time.Sleep(time.Duration((rand.Intn(4))) * time.Duration(time.Second))
		fmt.Printf("Processor - task #%d\n", id)
		wg.Done()
	}
}
