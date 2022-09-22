package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/benchungus/api/db"
	"github.com/benchungus/api/models"

	"github.com/benchungus/api/utils"
	"github.com/gin-gonic/gin"
)

//handles GET request from /monkeys endpoint, returns all monkeys in db
func GetMonkeys(c *gin.Context) {
	utils.InitializeLogger()
	utils.Logger.Info("grabbing all monkeys")
	monkeys := db.GrabAllMonkeys()
	utils.Logger.Info("got all monkeys")
	c.IndentedJSON(http.StatusOK, monkeys)
}

//handles GET request from /monkeys/:name endpoint, returns named monkey
func GetMonkeyByName(c *gin.Context) {
	name := c.Param("name")
	utils.InitializeLogger()
	utils.Logger.Info("grabbing monkey named " + name)
	monkey := db.GrabNamedMonkey(name)
	utils.Logger.Info("got monkey named " + name)
	c.IndentedJSON(http.StatusOK, monkey)
}

//handles POST request from /monkeys endpoint, inserts new monkey
func NewMonkey(c *gin.Context) {
	utils.InitializeLogger()
	var newMonkey models.Monkey
	err := json.NewDecoder(c.Request.Body).Decode(&newMonkey)
	if err != nil {
		log.Panic(err)
	}
	utils.Logger.Info("inserting monkey named " + newMonkey.Name)
	db.InsertMonkey(newMonkey)
	utils.Logger.Info("inserted monkey named " + newMonkey.Name)
}

//handles POST request from /monkeys/:name endpoint, updates hobby of named monkey
func ChangeHobby(c *gin.Context) {
	utils.InitializeLogger()
	name := c.Param("name")
	var newMonkey models.Monkey
	err := json.NewDecoder(c.Request.Body).Decode(&newMonkey)
	if err != nil {
		log.Panic(err)
	}
	utils.Logger.Info("updating monkey named " + name)
	db.UpdateMonkey(name, newMonkey.Hobby)
	utils.Logger.Info("updated monkey named " + newMonkey.Name)
}

//handles DELETE request from /monkeys/:name endpoint, kills named monkey
func KillMonkey(c *gin.Context) {
	name := c.Param("name")
	utils.InitializeLogger()
	utils.Logger.Info("killing monkey named " + name)
	db.DeleteMonkey(name)
	utils.Logger.Info("killed monkey named " + name)
}
