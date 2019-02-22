package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xiaoheigou/how-to-code/07-go-in-action-currency/runner/runner"
)

const timeout = 6 * time.Second

func main() {

	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask(), createTask(), createTask())
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
func createTask() func(int) {
	return func(id int) {
		fmt.Printf("Processor - task #%d\n", id)
		time.Sleep(time.Duration(time.Second))
	}
}
