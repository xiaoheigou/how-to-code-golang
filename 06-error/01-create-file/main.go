package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("test.go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	r := strings.NewReader("package main\nfunc main()")
	io.Copy(file, r)
}
