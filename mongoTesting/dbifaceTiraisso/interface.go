package dbiface

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CollectionAPI a collection API interface
type CollectionAPI interface {
	InsertOne(ctx context.Context, document interface{},opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	// DeleteOne(ctx context.Context, filter interface{},opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	// Find(ctx context.Context, filter interface{},opts ...*options.FindOptions) (cur *mongo.Cursor, err error) 
}
//Qualquer estrutura ou qualquer coisa em golang que implementa o método InsertOne com a mesma assinatura pode ser considerada um CollectionAPI, pois implementa implicitamente a CollectionAPI interface