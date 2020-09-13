package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
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
// 插入多条数据
func main() {

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

	record := &Record{
		JobName:   "job779",
		Command:   "echo sorry",
		Err:       "",
		Content:   "sorry",
		TimePoint: TimePoint{StartTime:time.Now().Unix(),EndTime:time.Now().Unix()+10},
	}

	logArr := []interface{}{record,record}
	insertManyRes,err := collection.InsertMany(context.TODO(),logArr)

	if err != nil{
		fmt.Println(err)
		return
	}
	docIds := insertManyRes.InsertedIDs
	fmt.Println("last_id:",docIds)

}