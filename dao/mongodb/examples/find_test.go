package examples

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestOne(t *testing.T) {
	var err error
	var client *mongo.Client
	var client *mongo.Collection
	var ctx = context.Background()
	var doc bson.M

}
