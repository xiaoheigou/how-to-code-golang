package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin goroutines", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("hello from this one")
		fmt.Println("mid goroutines", runtime.NumGoroutine())
		wg.Done()
	}()
	go func() {

		wg.Done()
	}()

	wg.Wait()
	fmt.Println("after cpu", runtime.NumCPU())
	fmt.Println("after goroutines", runtime.NumGoroutine())

}
