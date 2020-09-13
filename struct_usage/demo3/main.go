package main

import "fmt"


//Address 地址结构体
type Address struct {
	Province string
	City     string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address //匿名字段
	//Address Address//匿名字段
}

func main() {

	var user2 User
	user2.Name = "小王"
	user2.Gender = "男"
	user2.Address.Province = "山东"    // 匿名字段默认使用类型名作为字段名
	user2.City = "威海"                // 匿名字段可以省略

	fmt.Printf("user2=%#v\n", user2)
}



