package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN its connection object for mongodb */
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:4kVfqWJUbO2LhN6o@cursos.croox.mongodb.net/cursos?authSource=admin&ssl=true")

/* ConnectDB function for connect BD */
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connection Successfuly")
	return client
}

/* ConnectionCheck validate ping into DB */
func ConnectionCheck() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err != nil
}

// func CheckUserExist(email) bool{
// 	return true
// }