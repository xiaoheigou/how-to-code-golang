package main

import (
	"fmt"
	"time"
)

func main() {
	c := boring("boring!")

	//***********METHOD1
	// for i := 0; i < 1111111111111; i++ {
	// 	fmt.Printf("You say %q\n", <-c)
	// }
	// fmt.Println("U are boring;I'm leaving.")

	//***********METHOD2
	// for v := range c {
	// 	fmt.Printf("You say %q\n", v)
	// }

	//***********METHOD3
	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("You say %q\n", v)
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(time.Second))
		}
		close(c)
	}()

	return c
}
