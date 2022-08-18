package mongodb

import "testing"

func TestInit(t *testing.T) {
	InitMongoDB()
	collection := Client.Database("").Collection("")
	collection.InsertOne()
}