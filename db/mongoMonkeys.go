package db

import (
	"context"
	"log"

	"github.com/benchungus/api/models"
	"github.com/benchungus/api/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GrabAllMonkeys() []models.Monkey {
	utils.InitializeLogger()
	client := ConnDB()
	monkeyCollection := client.Database("jungle").Collection("monkies")
	cursor, err := monkeyCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var monkeys []models.Monkey

	if err = cursor.All(context.TODO(), &monkeys); err != nil {
		log.Fatal(err)
	}

	return monkeys
}

func GrabNamedMonkey(name string) models.Monkey {
	utils.InitializeLogger()
	client := ConnDB()
	var monkey models.Monkey
	monkeyCollection := client.Database("jungle").Collection("monkies")
	err := monkeyCollection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&monkey)
	if err != nil {
		log.Panic(err)
	}

	return monkey
}

func InsertMonkey(newMonkey models.Monkey) {
	utils.InitializeLogger()
	client := ConnDB()
	monkeyCollection := client.Database("jungle").Collection("monkies")
	_, err := monkeyCollection.InsertOne(context.TODO(), newMonkey)
	if err != nil {
		log.Panic(err)
	}
}

func UpdateMonkey(name string, hobby string) {
	utils.InitializeLogger()
	client := ConnDB()
	monkeyCollection := client.Database("jungle").Collection("monkies")
	_, err := monkeyCollection.UpdateOne(context.TODO(), bson.M{"name": name}, bson.D{{"$set", bson.D{{"hobby", hobby}}}})
	if err != nil {
		log.Panic(err)
	}

}

func DeleteMonkey(name string) {
	utils.InitializeLogger()
	client := ConnDB()
	monkeyCollection := client.Database("jungle").Collection("monkies")
	_, err := monkeyCollection.DeleteOne(context.TODO(), bson.M{"name": name})
	if err != nil {
		log.Panic(err)
	}
}
