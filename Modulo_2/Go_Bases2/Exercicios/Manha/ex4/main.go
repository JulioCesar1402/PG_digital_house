package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func calcMin(numbers ...float32) (float32, error) {
	if len(numbers) == 0 {
		return 0.0, errors.New("input is required")
	}
	minNumber := numbers[0]
	for _, number := range numbers {
		if minNumber > number {
			minNumber = number
		}
	}
	return minNumber, nil
}

func calcMax(numbers ...float32) (float32, error) {
	if len(numbers) == 0 {
		return 0.0, errors.New("input is required")
	}
	maxNumber := numbers[0]
	for _, number := range numbers {
		if maxNumber < number {
			maxNumber = number
		}
	}
	return maxNumber, nil
}

func calcAvg(numbers ...float32) (float32, error) {
	if len(numbers) == 0 {
		return 0.0, errors.New("input is required")
	}
	var medNumber float32
	for _, number := range numbers {
		medNumber += number
	}
	return medNumber / float32(len(numbers)), nil
}

func getFunc(t string) (func(values ...float32) (float32, error),
	error) {
	if t == minimum {
		return calcMin, nil
	}
	if t == average {
		return calcAvg, nil
	}
	if t == maximum {
		return calcMax, nil
	}
	return nil, errors.New("invalid function type")
}

func main() {
	avgFunc, _ := getFunc(average)
	maxFunc, _ := getFunc(maximum)
	minFunc, _ := getFunc(minimum)
	min, _ := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("min: %.2f\n", min)
	avg, _ := avgFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("avg: %.2f\n", avg)
	max, _ := maxFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Printf("max: %.2f\n", max)
}
