package main

import (
	"context"
	"fmt"
)

func main(){
	ctx ,cancel := context.WithCancel(context.Background())
	defer cancel()
// 消费者消费了5个之后调用cancel（） 通知生产者routine停止!!!
	for n := range gen(ctx){
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context)<-chan int {
		dst := make(chan int)
		n:=1
		go func(){
			for {
				select {
				case <-ctx.Done():
					return
				case dst<-n:
					n++
				}
			}
		}()
		return dst
}