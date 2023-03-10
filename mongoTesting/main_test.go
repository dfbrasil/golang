package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mockCollection struct{
}

func (m *mockCollection) InsertOne(ctx context.Context, document interface{},opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error){
	//fake implementation
	c := &mongo.InsertOneResult{}
	return c, nil
}

func TestInsertData(t *testing.T) {
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017")) //c é client
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	db := c.Database("tronics")
	col := db.Collection("products")
	res, err := insertData(col, User{"Daniel", "Brasil"})
	assert.Nil(t, err) //assert = afirmar
	assert.IsType(t, &mongo.InsertOneResult{}, res) //afirmar se res é um tipo InsertOneResult, res será um ponteiro para InsertOneResult
}