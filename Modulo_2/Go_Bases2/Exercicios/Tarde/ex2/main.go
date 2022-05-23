package main

import "fmt"

type loja struct {
	prod []produto
}

type produto struct {
	Tipo  string
	Nome  string
	Preco float64
}

type Ecommerce interface {
	Total() float64
	Adicionar(produto)
}

func novoProduto(tipo, nome string, preco float64) produto {
	return produto{tipo, nome, preco}
}

func (p produto) CalcularCusto() float64 {
	switch p.Tipo {
	case "Pequeno":
		return p.Preco
	case "Médio":
		return p.Preco + p.Preco * 3/100
	case "Grande":
		return p.Preco + p.Preco * 6/100 + 2500
	default:
		return 0
	}
}

func (l loja) Total() float64 {
	var sumTotal float64
	for _, t := range l.prod {
		sumTotal += t.CalcularCusto()
	}
	return sumTotal
}

func (l *loja) Adicionar(newProd produto) {
	l.prod = append(l.prod, newProd)
}

func novaLoja() Ecommerce {
	return &loja{}
}


func main() {
	lojaX := loja{}
	prod1 := novoProduto("Pequeno", "Pequeno", 10)
	prod2 := novoProduto("Médio", "Médio", 3000)
	prod3 := novoProduto("Grande", "Grande", 10000)

	lojaX.Adicionar(prod1)
	lojaX.Adicionar(prod2)
	lojaX.Adicionar(prod3)

	fmt.Println(lojaX)
	fmt.Println(lojaX.Total())
}