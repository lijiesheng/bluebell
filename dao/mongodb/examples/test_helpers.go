package examples

import (
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var (
	dbName = "argos"
	collectionName = "cars"
	collectionFavorites = "favorites"
	collectionExample = "examples"
)

var once sync.Once

func GetMongoClient(uri string) (*mongo.Client, error) {

}

