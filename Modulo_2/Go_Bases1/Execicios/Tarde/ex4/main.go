package main

import (
	"fmt"
	_ "time"
)

func main() {
	employee := "Benjamin"
	var employees = map[string]int{
		"Benjamin": 20, "Manuel": 26,
		"Brenda": 19, "Dario": 44, "Pedro": 30,
	}
	fmt.Printf("%s tem %d anos", employee, employees[employee])

	// month := time.Month(12)
	// fmt.Println(month)
}
