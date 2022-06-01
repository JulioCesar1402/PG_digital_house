package main

import "math/rand"

type client struct {
	File string
	Name string
	RG   int
	PhoneNumber int
	Address string
}

func generateID() int {
	return rand.Int()
}
