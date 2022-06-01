package main

import "fmt"

type Produto struct {
	Nome  string
	Preco float64
	Qty   int
}

type Usuario struct {
	Nome     string
	Produtos []Produto
}

func criarProduto(nome string, preco float64) Produto {
	return Produto{nome, preco, 1}
}

func adicionaProduto(user *Usuario, prod *Produto, qty int) {
	prod.Qty = qty
	user.Produtos = append(user.Produtos, *prod)
}

func deletarProduto(user *Usuario) {
	user.Produtos = user.Produtos[:0]
}

func main() {
	p1 := criarProduto("Celular", 2500)
	fmt.Printf("Novo Produto: %v\n", p1)
	fmt.Println("========================================")

	user := Usuario{Nome: "Gian"}
	adicionaProduto(&user, &p1, 5)
	fmt.Println("Usuario adicionou novo Produto:", user)
	fmt.Println("========================================")

	p2 := criarProduto("XBOX", 7500)
	p3 := criarProduto("gelo", 20)
	adicionaProduto(&user, &p2, 1)
	adicionaProduto(&user, &p3, 23)
	fmt.Println("Usuario:", user)
	fmt.Println("========================================")

	deletarProduto(&user)
	fmt.Println("Usuario com lista de Produtos vazia:", user)

}
