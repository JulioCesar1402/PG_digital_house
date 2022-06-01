package main

import "fmt"

type user struct {
	Nome string
	Sobrenome string
	Idade uint8
	Email string
	Senha string
}

func (u *user) mudarNome(nome string) {
	u.Nome = nome
}
func (u *user) mudarSobrenome(sobrenome string) {
	u.Sobrenome = sobrenome
}
func (u *user) mudarIdade(idade uint8) {
	u.Idade = idade
}
func (u *user) mudarEmail(email string) {
	u.Email = email
}
func (u *user) mudarSenha(senha string) {
	u.Senha = senha
}

func main() {
	p1 := user{"Gian", "Raphael", 22, "gian@gmail.com", "senha123"}
	fmt.Println(p1)
	p1.mudarNome("Gian Raphael")
	p1.mudarSobrenome("Martins")
	p1.mudarIdade(23)
	p1.mudarEmail("gian@hotmail.com")
	p1.mudarSenha("senha321")
	fmt.Println(p1)
}