package main

import (
	"context"
	"fmt"
	"log"
	"mongoTesting/dbiface"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
//User a dummy user
type User struct{
	FirstName string `bson:"first_name"`
	LastName string `bson:"last_name"`
}

//actual collection *mongo.Collection
//implements CollectionAPI
//Calls InsertOne will work on actual collection
//fake collection called mockCollection que também implementa CollectionAPI, O método InsertOne também terá uma implementação Fake, quanod invocada will work on fake collection


func insertData(collection dbiface.CollectionAPI, user User) (*mongo.InsertOneResult, error) { //trocou aqui *mongo.Collection por dbiface.CollectionAPI, trocou pela interface que tem a mesma assinatura
	res, err := collection.InsertOne(context.Background(), user) //essa função insertOne será testada
	if err != nil {
		return res, err
	}
	return res, nil
}

func main(){
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017")) //c é client
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	db := c.Database("tronics")
	col := db.Collection("products")
	res, err := insertData(col, User{"Daniel", "Brasil"})
	log.Println(res, err)
}
