package main

import (
	"github.com/benchungus/api/handler"
	"github.com/benchungus/api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	//load all env vars
	utils.GetEnvVar()

	//create router
	router := gin.Default()
	router.SetTrustedProxies(nil)

	//set up endpoints
	router.GET("/monkeys", handler.GetMonkeys)
	router.GET("/monkeys/:name", handler.GetMonkeyByName)
	router.POST("/monkeys", handler.NewMonkey)
	router.POST("/monkeys/:name", handler.ChangeHobby)
	router.DELETE("/monkeys/:name", handler.KillMonkey)

	//establish connection
	router.Run("localhost:8080")
}
