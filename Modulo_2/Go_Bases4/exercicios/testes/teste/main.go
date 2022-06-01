package main

import (
	"fmt"
	"os"
)

// estrutura personalizada de error
type myCustomError struct {
	status int
	msg    string
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

// aplicando o metodo de erro personalizado
func myCustomErrorTest(status int) (int, error) {
	if status >= 400 {
		return 500, &myCustomError{
			status: status,
			msg:    "Algo deu errado",
		}
	}
	return 200, nil
}

func main() {
	status, err := myCustomErrorTest(400)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Status %d, Funciona!\n", status)
}
