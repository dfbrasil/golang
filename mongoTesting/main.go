package main

import (
	"context"
	"fmt"
	// "log"
	"mongoTesting/dbiface"

	// "go.mongodb.org/mongo-driver/bson"
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

// func findData(collection dbiface.CollectionAPI) ([]User, error){
// 	var users []User
// 	ctx := context.Background()
// 	cur, err := collection.Find(ctx, bson.M{})// bson.M{} sem nada pois não tem nenhum critério de busca
// 	if err != nil{
// 		fmt.Printf("find error: %v\n", err)
// 		return users, err
// 	}
// 	fmt.Printf("cursor: %v\n", cur.Current)
// 	err = cur.All(ctx, &users)
// 	if err != nil {
// 		return users, err
// 	}
// 	return users, nil
// }

func main(){
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017")) //c é client
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	db := c.Database("tronics")
	col := db.Collection("products")
	res, err := insertData(col, User{"Daniel", "Brasil"})
	if err != nil {
		fmt.Printf("insert failure: %v\n", err)
	}
	fmt.Println(res)
	// users, err := findData(col)
	// if err != nil {
	// 	fmt.Printf("insert failure: %v\n", err)
	// }
	// fmt.Println(users)
}
