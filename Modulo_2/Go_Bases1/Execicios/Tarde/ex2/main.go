package main

import "fmt"

func main() {
	var (
		// age uint8 = 22
		age uint8 = 23
		employed bool = true
		timeWorked uint8 = 2
		// wage float32 = 100000.00
		wage float32 = 100000.01
	)
	switch {
	case age > 22 && employed && timeWorked > 1:
		if wage > 100000.00 {
			fmt.Println("Empréstimo concedido sem juros!")
		} else {
			fmt.Println("Empréstimo concedido com juros!")
		}
	default:
		fmt.Println("Empréstimo não concedido")
	}
}