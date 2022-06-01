package main

import (
	"fmt"
	"sync"
)

type Produtos struct {
	Nome  string
	Preco float64
	Qty   uint
}

type Servicos struct {
	Nome        string
	Preco       float64
	MinutosTrab uint
}

type Manutencao struct {
	Nome  string
	Preco float64
}

var soma float64
var wait sync.WaitGroup

func SomarProdutos(lista []Produtos) {
	for _, prod := range lista {
		fmt.Println("Soma produtos")
		soma += prod.Preco * float64(prod.Qty)
	}
	wait.Done()
}

func SomarServicos(lista []Servicos) {
	for _, serv := range lista {
		fmt.Println("Soma servicos")
		tmpServ := serv.MinutosTrab
		if tmpServ < 30 {
			soma += serv.Preco * 30
		} else {
			soma += serv.Preco * float64(tmpServ)
		}
	}
	wait.Done()
}

func SomarManutencao(lista []Manutencao) {
	var soma float64
	for _, manut := range lista {
		fmt.Println("Soma manutencao")
		soma += manut.Preco
	}
	wait.Done()
}

func main() {
	listaDeProdutos := []Produtos{
		{"t", 5, 2},
		{"t", 5, 2},
	}
	listaDeServicos := []Servicos{
		{"t", 3, 10},
		{"t", 1, 40},
	}
	listaDeManutencao := []Manutencao{
		{"t", 3},
		{"t", 1},
	}

	wait.Add(3)
	go SomarProdutos(listaDeProdutos)
	go SomarServicos(listaDeServicos)
	go SomarManutencao(listaDeManutencao)
	wait.Wait()
	fmt.Println(soma)
}
