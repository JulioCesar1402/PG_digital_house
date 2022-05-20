package main

import (
	"fmt"
)



func main() {
	employees := []map[string]any{
		{
			"Nome":      "Julio",
			"Categoria": "C",
			"Horas":     162,
		},
		{
			"Nome":      "Pedro",
			"Categoria": "B",
			"Horas":     176,
		},
		{
			"Nome":      "Cesar",
			"Categoria": "A",
			"Horas":     172,
		},
	}
	for _, employee := range employees {
		categoria := employee["Categoria"].(string)
		horasTrabalhadas := employee["Horas"].(int)
		fmt.Println(categoria)
		switch categoria {
		case "C":
			salario := calcSalary(horasTrabalhadas, 1000, 1, 100)
			fmt.Println(salario)
		case "B":
			salario := calcSalary(horasTrabalhadas, 1500, 160, 20)
			fmt.Println(salario)
		case "A":
			salario := calcSalary(horasTrabalhadas, 3000, 160, 50)
			fmt.Println(salario)
		}
	}
}

func calcSalary(horasTrabalhadas, salarioBase, horasBonus, taxa int) int {
	salario := horasTrabalhadas * salarioBase
	if horasTrabalhadas > horasBonus {
		salario += salario * taxa / 100
	}
	return salario
}