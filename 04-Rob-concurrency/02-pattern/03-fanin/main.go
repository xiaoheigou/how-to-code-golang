package main
import (
	"fmt"
	"math/rand"
	"time"
)
// zhou and bing are completely independent !
func fanIn(input1,input2 <-chan string)<-chan string{
	c := make(chan string)
	go func(){
		for {
			c<- <-input1
		}
	}()
	go func(){
		for {
			c<- <-input2
		}
	}()

	return c
}

func main(){
	c := fanIn(boring("zhou"),boring("bing"))

	for i := 0 ;i<5;i++{
		fmt.Printf("You say %q\n",<-c)
	}
	fmt.Println("U are boring;I'm leaving.")

}

func boring(msg string)<-chan string{
	c := make(chan string)
	go func(){
		for i :=0;;i++{
			c<-fmt.Sprintf("%s %d",msg,i)
			time.Sleep(time.Duration(rand.Intn(3))*time.Second)
		}
	}()
	return c 
}