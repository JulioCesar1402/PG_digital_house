package main

import (
	"fmt"
)

func main() {
	employees := []map[string]any{
		{
			"Nome":    "Julio",
			"Salario": 50000,
		},
		{
			"Nome":    "Gian",
			"Salario": 150000,
		},
		{
			"Nome":    "Roberto",
			"Salario": 50000,
		},
		{
			"Nome":    "Roberto",
			"Salario": 150000,
		},
	}
	verifyTax(employees)
}

func verifyTax(employees []map[string]any) {
	for _, employee := range employees {
		nome := employee["Nome"].(string)
		salario := employee["Salario"].(int)
		if salario <= 50000 {
			calculateTax(nome, salario, 100)
		} else if salario <= 150000 {
			calculateTax(nome, salario, 17)
		} else {
			calculateTax(nome, salario, 27)
		}
	}
}

func calculateTax(nome string, salario int, tax int) {
	fmt.Println("Funcionario:", nome)
	fmt.Println("Salário:", salario)
	desconto := salario * tax / 100
	fmt.Println("Desconto de:", desconto)
	fmt.Println("Salário com desconto:", salario-desconto)
	fmt.Println("==================================")
}
