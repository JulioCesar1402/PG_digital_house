package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	Id    int    `json:"id" binding:"-"`
	Color string `json:"color" binding:"required"`
}

var users []request
var lastId int

func AddColor(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != "123456" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "token inválido",
		})
		return
	}
	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	lastId++
	req.Id = lastId
	users = append(users, req)
	ctx.JSON(http.StatusCreated, req)
}

func main() {
	router := gin.Default()

	router.POST("/", AddColor)
	router.GET("/", GetAll())

	router.Run()
}

func GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}
