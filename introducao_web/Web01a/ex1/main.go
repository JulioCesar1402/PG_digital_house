package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type Products struct {
	Id        uint
	Name      string
	Color     string
	Price     float64
	Stock     int
	Code      string
	Published bool
	Date      string
}

var listOfProducts = []Products{
	{
		Id:        2,
		Name:      "MacBook",
		Color:     "White",
		Price:     10000,
		Stock:     50,
		Code:      "W#$ED%R%",
		Published: true,
		Date:      time.Now().Format("January 02, 2006"),
	},
	{
		Id:        3,
		Name:      "Fifine",
		Color:     "Black",
		Price:     270,
		Stock:     320,
		Code:      "u889&uy",
		Published: false,
		Date:      time.Now().Format("January 02, 2006"),
	},
}

type User struct {
	Id        uint
	FirstName string
	LastName  string
	Email     string
	Active    bool
	Date      string
}

var listOfUsers = []User{
	{
		Id:        23,
		FirstName: "Julio",
		LastName:  "Cesar",
		Email:     "julio.cesar@email.com",
		Active:    false,
		Date:      time.Now().Format("January 02, 2006"),
	},
	{
		Id:        5,
		FirstName: "Gian",
		LastName:  "Raphael",
		Email:     "gian.raphael@email.com",
		Active:    true,
		Date:      time.Now().Format("January 02, 2006"),
	},
}

type Transition struct {
	Id       uint
	Code     string
	Coin     string
	Value    float64
	Issuer   string
	Receiver string
	Date     string
}

var listOfTransition = []Transition{
	{
		Id:       34,
		Code:     "&$HNUSDN*&",
		Coin:     "USD",
		Value:    2900.3,
		Issuer:   "Julio",
		Receiver: "Gian",
		Date:     time.Now().Format("January 02, 2006"),
	},
}

func CreateJsonFile[T any](subject string, content T) {
	jsonData, err := json.Marshal(listOfProducts)
	fmt.Println(string(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	file := fmt.Sprintf("%s.json", subject)

	_ = ioutil.WriteFile(file, jsonData, 0644)
}

func main() {
	CreateJsonFile("products", listOfProducts)
	CreateJsonFile("users", listOfUsers)
	CreateJsonFile("transitions", listOfTransition)
}
