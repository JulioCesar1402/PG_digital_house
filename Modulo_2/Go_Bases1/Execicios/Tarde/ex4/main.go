package main

import (
	"fmt"
)

func main() {
	// Part 1
	employee := "Benjamin"
	var employees = map[string]int{
		"Benjamin": 20, "Manuel": 26,
		"Brenda": 19, "Dario": 44, "Pedro": 30,
	}
	fmt.Printf("%s tem %d anos\n", employee, employees[employee])

	// Part 2
	var countEmployees int
	for _, age := range employees {
		if age > 21 {
			countEmployees += 1
		}
	}
	fmt.Printf("%d funcionários têm mais de 21 anos\n", countEmployees)

	// Part 3
	employees["Federico"] = 25
	fmt.Printf("Federico tem %d anos\n", employees["Federico"])

	// Part 4
	fmt.Println("With Pedro:")
	listEmployees(employees)

	delete(employees, "Pedro")

	fmt.Println("Without Pedro:")
	listEmployees(employees)
}

func listEmployees(listOfEmployees map[string]int) {
	var employeesName []string
	for name := range listOfEmployees {
		employeesName = append(employeesName, name)
	}
	fmt.Println(employeesName)
}