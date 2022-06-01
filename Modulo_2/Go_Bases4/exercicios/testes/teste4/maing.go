package main

import (
	"errors"
	"fmt"
)

type myError struct {
	msg string
	x   string
}

func (e *myError) Error() string {
	return fmt.Sprintf("ocorreu um erro: %s, %s", e.msg, e.x)
}

func main() {
	e := &myError{"novo error", "404"}

	var err *myError

	isMyError := errors.As(e, &err)

	fmt.Println(isMyError)
}
