package main
import (
	"fmt"
	"time"
	"math/rand"
)
func main(){
	c := boring("zhou!")
	d := boring("bing")
	for i := 0 ;i<5;i++{
		fmt.Printf("You say %q\n",<-c)
		fmt.Printf("You say %q\n",<-d)
		// lock step
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