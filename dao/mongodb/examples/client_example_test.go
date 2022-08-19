package examples

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

const Cc_session_records = "cc_session_records"
const Cc_sheets = "cc_sheets"
const Cus_customers = "cus_customers"


// 来自 go.mongodb.org/mongo-driver/mongo 的 client_examples_test

type Student struct {     // _id 不用带入
	Name string
	Age int
	Tags []string
}

// 普通连接
func Test_ExampleClient(t *testing.T) {
	client , err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://192.168.6.249:27017"))
	if err != nil {
		fmt.Printf("mongodb connect failed, err : %s\n", err)
		return
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			fmt.Printf("mongodb disconnect failed, err : %s\n", err)
			return
		}
	}()
	collection := client.Database("erp").Collection("Student")

	findOptions := options.Find()
	findOptions.SetLimit(100)
	findOptions.SetSkip(20)
	findOptions.SetProjection(bson.M{"name" : 0})
	findOptions.SetSort(bson.D{{Key : "name" , Value: 1}})
	filter := bson.M{"name" : bson.M{"$in" : []string{"world", "world1"}}}
	distinct, err := collection.Distinct(context.TODO(), "name", filter)
	fmt.Println(len(distinct))
}


