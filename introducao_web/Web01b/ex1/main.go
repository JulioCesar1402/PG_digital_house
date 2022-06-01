package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Name  string
	Color string
	Price float64
}

var listOfProducts = []Product{
	{
		Name:  "Nintendo Switch",
		Color: "black",
		Price: 2500,
	},
	{
		Name:  "Smartwatch",
		Color: "white",
		Price: 83.21,
	},
	{
		Name:  "Microfone",
		Color: "black",
		Price: 132.11,
	},
	{
		Name:  "Livro",
		Color: "red",
		Price: 420.41,
	},
}

func main() {
	router := gin.Default()
	router.GET("/products", func(c *gin.Context) {
		colorQuery := c.Query("color")

		var searchForQuery []Product

		for _, prod := range listOfProducts {
			if prod.Color == colorQuery {
				searchForQuery = append(searchForQuery, prod)
			}
		}

		if len(searchForQuery) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"products": "Not found",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"products": searchForQuery,
			})
		}
	})
	router.Run()
}
