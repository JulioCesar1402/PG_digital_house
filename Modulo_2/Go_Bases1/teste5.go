package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c Circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// func detalhe(c Circulo) {
// 	fmt.Println(c)
// 	fmt.Println(c.area())
// 	fmt.Println(c.perimetro())
// }

// func main() {
// 	c := Circulo{raio: 5}
// 	detalhe(c)
// }

type geometria interface {
	area() float64
	perimetro() float64
}

type triangulo struct {
	largura, altura float64
}
func (t triangulo) area() float64 {
	return t.altura * t.largura
}
func (t triangulo) perimetro() float64 {
	return 2 * t.largura + t.altura
}

func detalhe(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

// func main() {
// 	t := triangulo{largura: 3, altura: 4}
// 	c := Circulo{raio: 5}
// 	detalhe(t)
// 	detalhe(c)
// }


const (
	tipoTriangulo = "triangulo"
	tipoCirculo = "circulo"
)

func novaFigura(geoTipo string, values ...float64) geometria {
	switch geoTipo {
	case tipoTriangulo:
		return triangulo{largura: values[0], altura: values[1]}
	case tipoCirculo:
		return Circulo{raio: values[0]}
	}
	return nil
}

func main() {
	t := novaFigura(tipoTriangulo, 3, 4)
	c := novaFigura(tipoCirculo, 5)
	detalhe(t)
	detalhe(c)
}
