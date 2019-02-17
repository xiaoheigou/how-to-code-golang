package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	// c := make(chan int ,10[5])  This is going to fail , because the channel doesnt block, quit channel will get 8 before all values in chan c get out !
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("get value sent ", i)
		}

		quit <- 8
	}()
	for {
		select {
		case v := <-c:
			time.Sleep(time.Second)
			fmt.Println(v)
		case v := <-quit:
			fmt.Println("quit value", v)
			return
		}
	}

}

/*get value sent  0
get value sent  1
get value sent  2
get value sent  3
get value sent  4
get value sent  5
0
get value sent  6
1
get value sent  7
2
get value sent  8
3
get value sent  9
4
5
6
quit value 8*/
