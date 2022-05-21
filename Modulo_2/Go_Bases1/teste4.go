package main

import "fmt"

type Veiculo struct {
	km   float64
	hora float64
}

func (v Veiculo) detalhe() {
	fmt.Printf("km:\t%f\nhora:\t%f\n", v.km, v.hora)
}

type Automovel struct {
	v Veiculo
}

func (a *Automovel) Correr(minutos int) {
	a.v.hora = float64(minutos) / 60
	a.v.km = a.v.hora * 100
}

func (a *Automovel) Detalhe() {
	fmt.Println("\nV:\tAutomovel")
	a.v.detalhe()
}

type Moto struct {
	v Veiculo
}

func (m *Moto) Correr(minutos int) {
	m.v.hora = float64(minutos) / 60
	m.v.km = m.v.hora * 80
}

func (m *Moto) Detalhe() {
	fmt.Println("\nV:\tMoto")
	m.v.detalhe()
}

func main() {
	automovel := Automovel{}
	automovel.Correr(360)
	automovel.Detalhe()

	moto := Moto{}
	moto.Correr(360)
	moto.Detalhe()
}