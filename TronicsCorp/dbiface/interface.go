package dbiface

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	// CollectionAPI is an interface for the mongo.Collection type
	CollectionAPI interface {
		InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
		Find(ctx context.Context, filter interface{},opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
		FindOne(ctx context.Context, filter interface{},ts ...*options.FindOneOptions) *mongo.SingleResult
		UpdateOne(ctx context.Context, filter interface{}, update interface{},opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	}
)