package main

import (
	"fmt"
	"math"
)

type Circle struct {
	raio float64
}

func (c Circle) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c Circle) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func (c *Circle) setRaio(r float64) {
	c.raio = r
}

func main() {
	c := Circle{raio: 5}
	fmt.Println(c.area())
	fmt.Println(c.perimetro())
	c.setRaio(10)
	fmt.Println(c.area())
	fmt.Println(c.perimetro())
}
