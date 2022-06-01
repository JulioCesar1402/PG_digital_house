package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.Open("customers.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}
