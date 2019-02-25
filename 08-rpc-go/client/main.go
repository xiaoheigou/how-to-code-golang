package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item
	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("connection errror:", err)
	}
	a := Item{"First", "A first item"}
	client.Call("API.AddItem", a, &reply)
	fmt.Println(reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println(db)
}
