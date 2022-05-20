package main

import "fmt"

type Students struct {
	Nome           string
	Sobrenome      string
	RG             int64
	DataSeAdmissao string
}

func (s Students) Detalhes(){
	fmt.Println("Nome:", s.Nome)
	fmt.Println("Sobrenome:", s.Sobrenome)
	fmt.Println("RG:", s.RG)
	fmt.Println("DataSeAdmissao:", s.DataSeAdmissao)
}

func main() {
	list := []Students{}

	list = append(list, Students{"Julio","Cesar", 5, "14/02"})
	list = append(list, Students{"Gian","Raphael", 7, "23/07"})

	for _, student := range list {
		student.Detalhes()
		fmt.Println("===========================")

	}

}