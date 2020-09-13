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

type FindByJobName struct {
	JobName string `bson:"job_name"`
}
// 查找
func main() {
	// 定义一个切片用来存储查询结果
	var results []*Record
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


	//cond := &FindByJobName{JobName:"job779"}

	//filter := bson.D{{"job_name","job779"}}
	//查询过滤 翻页
	//cur,err := collection.Find(context.TODO(),cond,findOptions)
	cur,err := collection.Find(context.TODO(),bson.D{{}},findOptions)


	//延迟释放
	defer cur.Close(context.TODO())
	if err != nil{
		fmt.Println(err)
		return
	}

	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(context.TODO()) {
		var elem Record

		//反序列化
		if err = cur.Decode(&elem);err != nil{
			fmt.Println(err)
			return
		}

		results = append(results, &elem)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found multiple documents (array of pointers): %#v\n", results)



}