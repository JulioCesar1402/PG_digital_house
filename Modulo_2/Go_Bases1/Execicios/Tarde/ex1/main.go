package main

import "fmt"

func main() {
	word := "Julio"
	fmt.Println("Palavra:", word)
	fmt.Println("Quantidade de letras:",len(word))
	fmt.Println("Letras:")
	for _, letter := range word {
		fmt.Printf("%c\n", letter)
	}
}