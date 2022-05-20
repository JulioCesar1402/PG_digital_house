package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func dogFunc(quantity int) int {
	return quantity * 10000
}

func catFunc(quantity int) int {
	return quantity * 5000
}

func hamsterFunc(quantity int) int {
	return quantity * 250
}

func tarantulaFunc(quantity int) int {
	return quantity * 150
}

func getCount(animal string) (func(qty int) int, error) {
	switch animal {
	case dog:
		return dogFunc, nil
	case cat:
		return catFunc, nil
	case hamster:
		return hamsterFunc, nil
	case tarantula:
		return tarantulaFunc, nil
	}
	return nil, errors.New("invalid function type")
}

func main() {
	dfunc, _ := getCount(dog)
	fmt.Printf("quandidade de alimento do(s) cachorro(s) (em gramas):%d gramas\n", dfunc(5))
	cfunc, _ := getCount(cat)
	fmt.Printf("quandidade de alimento do(s) gato(s) (em gramas): %dgramas\n", cfunc(8))
	_, err := getCount("invalid animal")
	if err != nil {
		fmt.Println("animal inválido")
	}
}
