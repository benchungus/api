package main

import (
	"github.com/benchungus/api/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/monkeys", handler.GetMonkeys)
	router.Run("localhost:8080")

}
