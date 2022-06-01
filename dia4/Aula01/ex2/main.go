package main

import (
	"errors"
	"fmt"
)

func VerifySalary(salary int) (int, error) {
	if salary < 15000 {
		return 500, errors.New("o salário digitado não está dentro do valor mínimo")
	}
	return 200, nil
}

func main()  {
	var salary int = 15000
	status, err := VerifySalary(salary)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(status)
}