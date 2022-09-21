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
	router.Run("localhost:8080")
}
