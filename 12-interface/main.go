package main

import (
	"flag"
	"fmt"
	"image/color"

	"github.com/campoy/tools/flags"
)

type person struct {
	val int
	sayi
}

type sayi interface {
	say()
}

func (p *person) run() {
	fmt.Println("i am running")
}

type me struct {
}

func (m *me) say() {
	fmt.Println("heheh")
}

type run interface {
	run()
}

func main() {
	var b color.Color
	flags.HexColorVar(&b, "ddd", color.Black, "fasdfasdfsa")
	flag.Parse()
	fmt.Printf("bbbbb------%v", b)

	m := &me{}
	p := new(sayi)
	hhh(p, m)
	fmt.Println(*p)

}
func hhh(c *sayi, value sayi) {
	p := &person{0, value}
	*c = p
	p.val = 9
	runner(p)

}
func runner(p run) {

	fmt.Println("runr")
}
