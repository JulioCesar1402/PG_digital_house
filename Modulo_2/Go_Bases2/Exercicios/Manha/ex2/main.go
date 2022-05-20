package main

import (
	"errors"
	"fmt"
)

func main() {
	students := []map[string]any{
		{
			"Nome":  "Julio",
			"Notas": []int{1, 2, 3},
		},
		{
			"Nome":    "Gian",
			"Notas": []int{1, 2},
		},
		{
			"Nome":    "Roberto",
			"Notas": []int{1, 2, -3},
		},
	}
	for _, student := range students {
		grades := student["Notas"].([]int)
		med, err := calcMed(grades...)
		if err == nil {
			fmt.Println(med)
		} else {
			fmt.Println(err)
		}
	}
}

func calcMed(grades ...int) (float32, error) {
	var sumOfGrades int
	for _, num := range grades {
		if num < 0 {
			return 0, errors.New("um dos números digitados é negativo")
		}
		sumOfGrades += num
	}
	qtyGrades := float32(len(grades))
	return float32(sumOfGrades) / qtyGrades, nil
}
