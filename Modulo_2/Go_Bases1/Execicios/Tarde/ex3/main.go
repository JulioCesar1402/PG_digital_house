package main

import (
	"fmt"
	_ "time"
)

func main() {
	month := 2
	var months = []string{
		"January", "February", "March", "April",
		"May", "June", "July", "August",
		"September", "October", "November", "December",
	}
	fmt.Println(months[month - 1])

	// month := time.Month(12)
	// fmt.Println(month)
}