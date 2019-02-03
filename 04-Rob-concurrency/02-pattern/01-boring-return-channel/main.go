package main
import (
	"fmt"
	"time"
)
func main(){
	c := boring("boring!")
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
			time.Sleep(time.Duration(time.Second))
		}
	}()
	return c 
}