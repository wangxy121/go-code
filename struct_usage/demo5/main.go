package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id int			`json:"id"`
	Gender string	`json:"gender"`
	Name string		`json:"name"`
}

type Class struct {
	Title    string		`json:"title"`
	Student []*Student	`json:"student"`
}
func main() {
	c := &Class{
		Title:   "一班",
		Student: make([]*Student,0,20),
	}

	for i := 0; i < 3; i++ {
		stu := &Student{
			Id:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu%02d", i),
		}
		c.Student = append(c.Student,stu)
	}

	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)


	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"一班","Student":[{"Id":0,"Gender":"男","Name":"stu00"},{"Id":1,"Gender":"男","Name":"stu01"},{"Id":2,"Gender":"男","Name":"stu02"}]}`

	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}
