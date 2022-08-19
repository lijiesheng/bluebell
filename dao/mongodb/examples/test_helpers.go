package examples

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	return getMongoClientByURI(uri)
}

// 通过 uri 连接 mongodb
func getMongoClientByURI(uri string) (*mongo.Client, error) {
	var err error
	var client *mongo.Client
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		fmt.Printf("mongo connect failed, err : %v", err)
		return nil, err
	}
	client.Ping(context.Background(), nil)
	return client,err
}

func getMongoClient() (*mongo.Client, error) {
	uri := "mongodb://192.168.6.249/erp?replicaSet=replset"
	return getMongoClientByURI(uri)
}

//// todo 不知道有啥用
//func seedCarsData(client *mongo.Client, database string) int64{
//	var err error
//	var count int64
//	collection := client.Database(database).Collection(collectionName)
//	filter := bson.D{{}}
//	if count, err = collection.CountDocuments(context.Background(), filter); err != nil {
//		fmt.Println("===>", err)
//		return 0
//	}
//	if count == 0 {
//		f := New
//	}
//}



