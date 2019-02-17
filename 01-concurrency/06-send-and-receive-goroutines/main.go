package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
			time.Sleep(time.Second)
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v, "-----get the value from the channel c")
	}
}
