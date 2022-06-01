package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Name string
	Price float64
}

var listOfProducts = []Product {
	{
		Name: "Nintendo Switch",
		Price: 2500,
	},
	{
		Name: "Smartwatch",
		Price: 83.21,
	},
}

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"products": listOfProducts,
	})
}

func main() {
	router := gin.Default()
	router.GET("/products", GetAll)
	router.Run()
}
