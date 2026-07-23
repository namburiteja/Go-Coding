package main

type Employee struct {
	Name   string   `bson:"name"`
	Age    int      `bson:"age"`
	City   string   `bson:"city"`
	Salary int      `bson:"salary"`
	Skills []string `bson:"skills"`
}