package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	// go func(ctx context.Context) {
	// 	for {
	// 		time.Sleep(time.Second)
	// 		fmt.Println("------")

	// 	}
	// }(ctx)
	dosomejob(ctx)

}

func dosomejob(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(time.Second * 4):
		fmt.Println("time-out-five-second")
	}
}
