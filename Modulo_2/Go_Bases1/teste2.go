package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Pessoa struct {
	PrimeiroNome string `json:"primeiro_nome" bd:"primeiro_nome"`
	Sobrenome    string `json:"sobrenome" bd:"sobrenome"`
	Telefone     string `json:"telefone" bd:"telefone"`
	Endereco     string `json:"endereco" bd:"endereco"`
}

func main() {
	p := Pessoa{"Paula", "Monteiro", "42342342", "Rua Limoeiro"}
	meuJSON, err := json.Marshal(p)
	fmt.Println(string(meuJSON))
	fmt.Println(err)

	fmt.Println("==============================")

	pessoa := Pessoa{}
	pw := reflect.TypeOf(pessoa)
	fmt.Println("Type:", pw.Name())
	fmt.Println("Kind:", pw.Kind())

	fmt.Println("==============================")

	for i :=0; i < pw.NumField(); i+=1{
		field := pw.Field(i)
		tag := field.Tag.Get("bd")
		fmt.Println(tag)
	}
}
