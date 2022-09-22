package db

import (
	"context"
	"log"

	"github.com/benchungus/api/models"
	"github.com/benchungus/api/utils"
	"go.mongodb.org/mongo-driver/bson"
)

//searches for all monkeys in jungle db and returns as an array
func GrabAllMonkeys() []models.Monkey {
	utils.InitializeLogger()
	client := ConnDB()
	monkeyCollection := client.Database("jungle").Collection("monkies")

	//generating cursor for this search
	cursor, err := monkeyCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var monkeys []models.Monkey

	//adding all found monkeys to an array
	if err = cursor.All(context.TODO(), &monkeys); err != nil {
		log.Fatal(err)
	}

	return monkeys
}

//searches for single monkey by name and returns info
func GrabNamedMonkey(name string) models.Monkey {
	utils.InitializeLogger()
	client := ConnDB()
	var monkey models.Monkey
	monkeyCollection := client.Database("jungle").Collection("monkies")

	//find and decode into monkey
	err := monkeyCollection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&monkey)
	if err != nil {
		log.Panic(err)
	}

	return monkey
}

//inserts new monkey into jungle db
func InsertMonkey(newMonkey models.Monkey) {
	utils.InitializeLogger()
	client := ConnDB()
	monkeyCollection := client.Database("jungle").Collection("monkies")
	_, err := monkeyCollection.InsertOne(context.TODO(), newMonkey)
	if err != nil {
		log.Panic(err)
	}
}

//updates a monkey's hobby by searching for name
func UpdateMonkey(name string, hobby string) {
	utils.InitializeLogger()
	client := ConnDB()
	monkeyCollection := client.Database("jungle").Collection("monkies")
	_, err := monkeyCollection.UpdateOne(context.TODO(), bson.M{"name": name}, bson.D{{"$set", bson.D{{"hobby", hobby}}}})
	if err != nil {
		log.Panic(err)
	}

}

//deletes a monkey from the database
func DeleteMonkey(name string) {
	utils.InitializeLogger()
	client := ConnDB()
	monkeyCollection := client.Database("jungle").Collection("monkies")
	_, err := monkeyCollection.DeleteOne(context.TODO(), bson.M{"name": name})
	if err != nil {
		log.Panic(err)
	}
}
