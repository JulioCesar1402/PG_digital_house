package userinfo

import "fmt"

func Info() {
	var (
		nome  = "Julio Cesar"
		idade = 19
	)
	fmt.Println("Nome:", nome, "-", "Idade:", idade)
}