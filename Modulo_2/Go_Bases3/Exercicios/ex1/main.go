package main

import (
	"fmt"
	"os"
)

type objs struct {
	ID    uint
	Price float64
	QTY   float64
}

// const (
// 	ID    uint    = 204
// 	Price float64 = 3576.0
// 	QTY   float64   = 137
// )

func main() {
	loja := []objs{}

	loja = append(loja, objs{111223, 30012.00, 1})
	loja = append(loja, objs{444321, 1000000.00, 4})
	loja = append(loja, objs{434321, 50.50, 1})

	for _, t := range loja {
		base, error2 := os.ReadFile("./MyFile.txt")
		teste := fmt.Sprintf("%v %25v %25v\n", "ID", "Preco", "Quantidade")
		if error2 != nil {
			base = []byte(teste)
		}
		numberData := fmt.Sprintf("%d %20.2f %20.2f\n", t.ID, t.Price, t.QTY)
		data := fmt.Sprintf("%v %v", string(base), numberData)
		text := []byte(data)

		err := os.WriteFile("./MyFile.txt", text, 0644)

		if err != nil {
			fmt.Println(str)
		}
	}
}
