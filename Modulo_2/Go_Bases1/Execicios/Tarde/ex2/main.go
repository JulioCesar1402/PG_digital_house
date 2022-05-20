package main

import "fmt"

// func main() {
// 	var (
// 		// age uint8 = 22
// 		age uint8 = 23
// 		employed bool = true
// 		timeWorked uint8 = 2
// 		// wage float32 = 100000.00
// 		wage float32 = 100000.01
// 	)
// 	switch {
// 	case age > 22 && employed && timeWorked > 1:
// 		if wage > 100000.00 {
// 			fmt.Println("Empréstimo concedido sem juros!")
// 		} else {
// 			fmt.Println("Empréstimo concedido com juros!")
// 		}
// 	default:
// 		fmt.Println("Empréstimo não concedido")
// 	}
// }

func main() {
	clientes := []map[string]interface{}{
		{
			"Nome":      "João",
			"Idade":     21,
			"Atividade": 1,
			"Salário":   100000,
		},
		{
			"Nome":      "José",
			"Idade":     25,
			"Atividade": 1,
			"Salário":   100000,
		},
		{
			"Nome":      "Carlos",
			"Idade":     27,
			"Atividade": 2,
			"Salário":   99000,
		},
		{
			"Nome":      "António",
			"Idade":     35,
			"Atividade": 5,
			"Salário":   150000,
		},
	}
	for _, cliente := range clientes {
		nome := cliente["Nome"]
		fmt.Println("Cliente:", nome)

		idade := cliente["Idade"].(int)
		if idade <= 22 {
			fmt.Println("Não possui empréstimo disponível (idade insuficiente)")
			continue
		}

		atividade := cliente["Atividade"].(int)
		if atividade <= 1 {
			fmt.Println("Não possui empréstimo disponível (atividade insuficiente)")
			continue
		}

		salario := cliente["Salário"].(int)
		if salario > 100000 {
			fmt.Println("Possui empréstimo disponível sem juros")
		} else {
			fmt.Println("Possui empréstimo disponível com juros")
		}
	}
}
