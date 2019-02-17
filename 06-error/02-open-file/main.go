package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("../01-create-file/test.go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	bs, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
