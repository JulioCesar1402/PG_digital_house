package main

import (
	"fmt"

	"github.com/juliocesar1402/ex-manha/pkg/evaluation"
	userinfo "github.com/juliocesar1402/ex-manha/pkg/userInfo"
	"github.com/juliocesar1402/ex-manha/pkg/weather"
)

func main() {
	// ex 1
	userinfo.Info()
	// ex 2
	weather.Weather()
	// ex 3
	evaluation.Exame()
	// ex 4
	var sobrenome string = "Silva"
	var idade uint8 = 25
	boolean := false
	var salario float32 = 4585.90
	var nome string = "Fellipe"

	fmt.Println(nome, sobrenome, idade, boolean, salario, nome)
}
