package handler

import (
	"net/http"

	"github.com/benchungus/api/db"
	//"github.com/benchungus/api/models"
	"github.com/benchungus/api/utils"
	"github.com/gin-gonic/gin"
)

func GetMonkeys(c *gin.Context) {
	utils.InitializeLogger()
	utils.Logger.Info("grabbing all monkeys")
	monkeys := db.GrabAllMonkeys()
	utils.Logger.Info("got all monkeys")
	c.IndentedJSON(http.StatusOK, monkeys)
}

func GetMonkeyByName(c *gin.Context) {
	name := c.Param("name")
	utils.InitializeLogger()
	utils.Logger.Info("grabbing monkey named " + name)
	monkey := db.GrabNamedMonkey(name)
	utils.Logger.Info("got monkey named " + name)
	c.IndentedJSON(http.StatusOK, monkey)
}
