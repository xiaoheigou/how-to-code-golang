package main

import (
	"fmt"
)

type MenuComponent interface {
	Name() string
	Description() string
	Price() float32
	Print()

	Add(MenuComponent)
	Remove(int)
	Child(int) MenuComponent
}

type MenuItem struct {
	name        string
	description string
	price       float32
}

func NewMenuItem(name, description string, price float32) MenuComponent {
	return &MenuItem{
		name:        name,
		description: description,
		price:       price,
	}
}

func (m *MenuItem) Name() string {
	return m.name
}

func (m *MenuItem) Description() string {
	return m.description
}

func (m *MenuItem) Price() float32 {
	return m.price
}

func (m *MenuItem) Print() {
	fmt.Printf("  %s, ￥%.2f\n", m.name, m.price)
	fmt.Printf("    -- %s\n", m.description)
}

func (m *MenuItem) Add(MenuComponent) {
	panic("not implement")
}

func (m *MenuItem) Remove(int) {
	panic("not implement")
}

func (m *MenuItem) Child(int) MenuComponent {
	panic("not implement")
}

type Menu struct {
	name        string
	description string
	children    []MenuComponent
}

func NewMenu(name, description string) MenuComponent {
	return &Menu{
		name:        name,
		description: description,
	}
}

func (m *Menu) Name() string {
	return m.name
}

func (m *Menu) Description() string {
	return m.description
}

func (m *Menu) Price() (price float32) {
	for _, v := range m.children {
		price += v.Price()
	}
	return
}

func (m *Menu) Print() {
	fmt.Printf("%s, %s, ￥%.2f\n", m.name, m.description, m.Price())
	fmt.Println("------------------------")
	for _, v := range m.children {
		v.Print()
	}
	fmt.Println()
}

func (m *Menu) Add(c MenuComponent) {
	m.children = append(m.children, c)
}

func (m *Menu) Remove(idx int) {
	m.children = append(m.children[:idx], m.children[idx+1:]...)
}

func (m *Menu) Child(idx int) MenuComponent {
	return m.children[idx]
}

func main() {
	menu1 := NewMenu("培根鸡腿燕麦堡套餐", "供应时间：09:15--22:44")
	menu1.Add(NewMenuItem("主食", "培根鸡腿燕麦堡1个", 11.5))
	menu1.Add(NewMenuItem("小吃", "玉米沙拉1份", 5.0))
	menu1.Add(NewMenuItem("饮料", "九珍果汁饮料1杯", 6.5))

	menu2 := NewMenu("奥尔良烤鸡腿饭套餐", "供应时间：09:15--22:44")
	menu2.Add(NewMenuItem("主食", "新奥尔良烤鸡腿饭1份", 15.0))
	menu2.Add(NewMenuItem("小吃", "新奥尔良烤翅2块", 11.0))
	menu2.Add(NewMenuItem("饮料", "芙蓉荟蔬汤1份", 4.5))

	all := NewMenu("超值午餐", "周一至周五有售")
	all.Add(menu1)
	all.Add(menu2)

	all.Print()
}
