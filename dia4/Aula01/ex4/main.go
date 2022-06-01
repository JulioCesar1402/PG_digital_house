package main

import (
	"errors"
	"fmt"
)

func calcSalary(hours int, value float64) (float64, error) {
	salaryPerMonth := float64(hours) * value
	switch {
	case hours < 80 || hours < 0:
		return 0, fmt.Errorf("error: %w", errors.New("o número de horas mensais é menor que 80 ou negativo"))
	case salaryPerMonth >= 20000:
		return salaryPerMonth - salaryPerMonth * 0.1, nil
	default:
		return salaryPerMonth, nil
	}
}

func main() {
	sa , err := calcSalary(10, 12)
	if err != nil {
		fmt.Println(errors.Unwrap(err))
		return
	}
	fmt.Println(sa)
}