package main

import (
	"github.com/benchungus/api/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/monkeys", handler.GetMonkeys)
	router.GET("/monkeys/:name", handler.GetMonkeyByName)
	router.POST("/monkeys", handler.NewMonkey)
	router.POST("/monkeys/:name", handler.ChangeHobby)
	router.DELETE("/monkeys/:name", handler.KillMonkey)
	router.Run("localhost:8080")
}
