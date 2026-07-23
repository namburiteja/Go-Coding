package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InsertEmployee(collection *mongo.Collection) {

	employee := Employee{
		Name:   "Rohit sai",
		Age:    22,
		City:   "Banglore",
		Salary: 70000,
		Skills: []string{"Frontend", "MySQL"},
	}

	result, err := collection.InsertOne(context.Background(), employee)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted ID:", result.InsertedID)
}

func FindOneEmployee(collection *mongo.Collection) {

	var emp Employee

	err := collection.FindOne(
		context.Background(),
		bson.M{"name": "Rohit sai"},
	).Decode(&emp)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nEmployee Found")
	fmt.Println(emp)
}

func FindAllEmployees(collection *mongo.Collection) {

	cursor, err := collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		log.Fatal(err)
	}

	var employees []Employee

	err = cursor.All(context.Background(), &employees)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nAll Employees")

	for _, emp := range employees {
		fmt.Println(emp)
	}
}

func UpdateEmployee(collection *mongo.Collection) {

	update := bson.M{
		"$set": bson.M{
			"salary": 50000,
		},
	}

	result, err := collection.UpdateOne(
		context.Background(),
		bson.M{"name": "Rohit sai"},
		update,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nModified:", result.ModifiedCount)
}

func DeleteEmployee(collection *mongo.Collection) {

	result, err := collection.DeleteOne(
		context.Background(),
		bson.M{"name": "Teja"},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted:", result.DeletedCount)
}
