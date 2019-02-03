package main
import (
	"fmt"
	"time"
	"math/rand"
)
//another fanin
func fanIn(input1 ,input2<-chan string)<-chan string{
	c := make(chan string)
	go func (){
		for {
			select {
			case s := <-input1:c<-s
			case s := <-input2:c<-s
			}
		}
	}()
	return c
}

//timeout 
func main(){
	c := boring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
			// timeout for each conversation,move the time.After outside the for loop
		case <-time.After(3*time.Second):
			fmt.Println("You're too show")
			return
		}
	}
}






func boring(msg string)<-chan string{
	c := make(chan string)
	go func(){
		for i :=0;;i++{
			c<-fmt.Sprintf("%s %d",msg,i)
			time.Sleep(time.Duration(rand.Intn(2))*time.Second)
		}
	}()
	return c 
}