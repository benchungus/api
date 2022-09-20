package handler

import (
	"net/http"

	"github.com/benchungus/api/models"
	"github.com/gin-gonic/gin"
)

var monkeys = []models.Monkey{
	{Name: "Ben", Type: "Golden Monkey"},
	{Name: "Caleb", Type: "Black Howler"},
	{Name: "Ryan", Type: "Mandrill"},
}

func GetMonkeys(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, monkeys)
}
