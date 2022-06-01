package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int
	Name string
}

var listOfUsers = []User{
	{
		Id:   234,
		Name: "Julio Cesar",
	},
	{
		Id:   400,
		Name: "Gian Raphael",
	},
}

func handleUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"user": "Invalid input",
		})
		log.Print(err.Error())
		return
	}

	for _, user := range listOfUsers {
		if user.Id == id {
			ctx.JSON(http.StatusOK, gin.H{
				"user": user,
			})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"user": "User not found",
	})
}

func main() {
	router := gin.Default()
	router.GET("/users/:id", handleUser)
	router.Run()
}
