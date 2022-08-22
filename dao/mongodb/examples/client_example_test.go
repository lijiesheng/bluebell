package examples

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
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
	var  stu1 Student
	collection := client.Database("erp").Collection("Student")
	// bson.D{{"$set", bson.D{{"email", "newemail@example.com"}}}}
	collection.FindOneAndUpdate(context.TODO(), bson.M{"_id": objectIDHex("62ff1b1d840c01f2575560c8")}, bson.M{"$set" : bson.M{"age" : 66}}).Decode(&stu1)


	//groupStage := bson.D{
	//	{"$group", bson.D{
	//		{"_id", "$name"},    // 根据 name 进行分组
	//		{"numTimes", bson.D{   // 每个 name 出现的次数，赋值给 numTimes
	//			{"$sum", 1},
	//		}},
	//	}},
	//	{"$match" : bson.D{
	//		{"numTimes" : bson.D{
	//			{"$gt" : 2},
	//		}},
	//	}},
	//}

	pipelint := mongo.Pipeline{
		{{Key : "$group", Value : bson.D{ {Key : "_id", Value : "$name"}, {Key : "numTimes", Value: bson.D{{Key: "$sum", Value : 1}}} }}},
		{{Key : "$match", Value : bson.D{{Key : "numTimes", Value: bson.D{{Key : "$gt", Value : 2}}}}}},
	}

	opt := options.Aggregate()
	cursor ,err := collection.Aggregate(context.TODO(),pipelint,opt)
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Printf(
			"name %v appears %v times\n",
			result["_id"],
			result["numTimes"])
	}

}

func MongoPipeline(str string) mongo.Pipeline {
	var pipeline = []bson.D{}
	str = strings.TrimSpace(str)
	if strings.Index(str, "[") != 0 {
		var doc bson.D
		bson.UnmarshalExtJSON([]byte(str), false, &doc)
		pipeline = append(pipeline, doc)
	} else {
		bson.UnmarshalExtJSON([]byte(str), false, &pipeline)
	}
	return pipeline
}


func objectIDHex(s string) primitive.ObjectID {
	oid, _ := primitive.ObjectIDFromHex(s)
	return oid
}
