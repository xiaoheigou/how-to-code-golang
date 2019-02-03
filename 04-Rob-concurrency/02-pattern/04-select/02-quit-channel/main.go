package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan string)
	ch := generator("Hi!", quit)
	for i := rand.Intn(13); i >= 0; i-- {
		fmt.Println(<-ch, i)
	}
	quit <- "Bye!"
	fmt.Printf("Generator says %s", <-quit)
}

func generator(msg string, quit chan string) <-chan string { // returns receive-only channel
	ch := make(chan string)
	go func() { // anonymous goroutine
		for {
			select {
			case ch <- fmt.Sprintf("%s", msg):
				// nothing
			case <-quit:
				quit <- "See you!"
				return
			}
		}
	}()
	return ch
}