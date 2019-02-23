package main

import (
	"log"
	"sync"
	"time"

	"github.com/xiaoheigou/how-to-code/07-go-in-action-currency/work/work"
)

var names = []string{
	"steve",
	"hhdhdh",
	"axing",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}
func main() {
	p := work.New(2)
	var wg sync.WaitGroup
	wg.Add(100 * len(names))
	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()

	p.ShutDown()
}
