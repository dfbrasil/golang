package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Product describes an eletronic product e.g. phone
type Product struct{
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Name string `json:"product_name" bson:"product_name"`
	Price int `json:"price" bson:"price"`
	Currency string `json:"currency" bson:"currency"`
	Quantity int `json:"quantity" bson:"quantity"`
	Discount int `json:"discount,omitempty" bson:"discount,omitempty"`
	Vendor string `json:"vendor" bson:"vendor"`
	Accessories []string `json:"accessories"`
	SkuID string `json:"sku_id" bson:"sku_id"`
}

var iphone10 = Product{
	ID: primitive.NewObjectID(),
	Name: "iphone10",
	Price: 900,
	Currency: "USD",
	Quantity: 40,
	Vendor: "apple",
	Accessories: []string{"charger","headsset","slotopener"},
	SkuID: "1234",
}


var trimmer = Product{
	ID: primitive.NewObjectID(),
	Name: "easy philips trimmer",
	Price: 120,
	Currency: "USD",
	Quantity: 300,
	Vendor: "Philips",
	Discount: 7,
	Accessories: []string{"charger","comb", "bladeset", "cleaning oil"},
	SkuID: "2345",
}

var speaker = Product{
	ID: primitive.NewObjectID(),
	Name: "speakers",
	Price: 300,
	Currency: "USD",
	Quantity: 25,
	Vendor: "Bosch",
	Discount: 4,
	Accessories: []string{"cables","remote"},
	SkuID: "4567",
}

func main(){
	
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil{
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil{
		fmt.Println(err)
	}
	db := client.Database("tronics")
	collection := db.Collection("products")

	//***********CREATE************

	//using structs
	res, err := collection.InsertOne(context.Background(), trimmer)
	fmt.Println("---- Insert Many using struct ----")
	fmt.Println(res.InsertedID)

	//using BSON.M
	// res, err := collection.InsertOne(context.Background(), bson.M{
	// 	"name":"daniel",
	// 	"sobrenome":"brasil",
	// 	"hobbies":bson.A{
	// 		"videogames",
	// 		"tv",
	// 		"futebol",
	// 	},
	// })
	// fmt.Println("---- Insert Many using bson.M ----")
	// fmt.Println(res.InsertedID)

	//using BSON.D
	// res, err := collection.InsertOne(context.Background(), bson.D{
	// 	"name":"daniel",
	// 	"sobrenome":"brasil",
	// 	"hobbies":bson.A{
	// 		"videogames",
	// 		"tv",
	// 		"futebol",
	// 	},
	// })
	// fmt.Println("---- Insert Many using bson.D ----")
	// fmt.Println(res.InsertedID)

	//Insert many documents at the same time
	resMany, err := collection.InsertMany(context.Background(),[]interface{}{iphone10,speaker})
	fmt.Println("---- Insert Many ----")
	fmt.Println(resMany.InsertedIDs)


	//******* READ *********

	//Equality operator using FindOne
	var findOne Product
	err = collection.FindOne(context.Background(), bson.M{"price":900}).Decode(&findOne)
	fmt.Println("---- Equality operator using FindOne ----")
	fmt.Println(findOne)

	//Comparsion operator using Find
	var find Product
	fmt.Println("---- Comparsion operator using Find ----")
	findCursor, err := collection.Find(context.Background(), bson.M{"price": bson.M{"$gt":100}}) //$gt greater than --- mongo operator
	for findCursor.Next(context.Background()) {
		err := findCursor.Decode(&find)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(find.Name)
	}

	//Logical operator using Find
	var findLogic Product
	logicFilter := bson.M{
		"$and":bson.A{
			bson.M{"price": bson.M{"$gt":100}},
			bson.M{"quantity": bson.M{"$gt":30}},
		},
	}
	fmt.Println("---- Logical operator using Find ----")
	findLogicRes, err := collection.Find(context.Background(), logicFilter)
	for findLogicRes.Next(context.Background()) {
		err := findLogicRes.Decode(&findLogic)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(findLogic.Name)
	}

	//Element operator using Find
	var findElement Product
	elementFilter := bson.M{
		"accessories": bson.M{"$exists": true},
	}
	fmt.Println("---- Elemental operator using Find ----")
	findElementRes, err := collection.Find(context.Background(), elementFilter)
	for findElementRes.Next(context.Background()) {
		err := findElementRes.Decode(&findElement)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(findElement.Name)
	}

	// Array operator using Find
	var findArray Product
	arrayFilter := bson.M{"accessories": bson.M{"$all": bson.A{"charger"}}}
	fmt.Println("---- Array operator using Find ----")
	findArrayRes, err := collection.Find(context.Background(), arrayFilter)
	for findArrayRes.Next(context.Background()){
		err := findArrayRes.Decode(&findArray)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(findArray.Name)
	}

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
}

