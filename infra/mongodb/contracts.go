package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -destination=./mocks.go -package=mongodb -source=./contracts.go

type IRepository interface {
	InsertOne(ctx context.Context, collection string, data any) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, collection string, data any, opts *options.FindOneOptions) *mongo.SingleResult
	Find(ctx context.Context, collection string, data any, opts *options.FindOptions) (*mongo.Cursor, error)
	UpdateOne(ctx context.Context, collection string, filter any, data any) (*mongo.UpdateResult, error)
}
