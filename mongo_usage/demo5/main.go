package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type TimePoint struct {
	StartTime int64	`bson:"start_time"`
	EndTime   int64	`bson:"end_time"`
}

type Record struct {
	JobName string `bson:"job_name"`
	Command string  `bson:"command"`
	Err string		`bson:"err"`
	Content string	`bson:"content"`

	TimePoint TimePoint	`bson:"time_point"`
}


// 查找
func main() {
	var result Record
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://root:123456@localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// 指定获取要操作的数据库和表
	collection := client.Database("my_db").Collection("my_collection")

	// 将选项传递给Find()
	findOptions := options.Find()
	findOptions.SetSkip(0)
	findOptions.SetLimit(2)


	filter := bson.D{{"job_name","job779"}}



	err = collection.FindOne(context.TODO(),filter).Decode(&result)


	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("Found a single document: %+v\n", result)


}