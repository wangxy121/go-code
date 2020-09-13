package main

import (
	"fmt"
)

/*type 类型名 struct {
	字段名 字段类型
	字段名 字段类型
	…
}*/

/*类型名：标识自定义结构体的名称，在同一个包内不能重复。
字段名：表示结构体字段名。结构体中的字段名必须唯一。
字段类型：表示结构体字段的具体类型。*/

type person struct {
	name string
	city string
	age   int8
}
//同样类型的字段也可以写在一行
type person1 struct {
	name,city string
	age int8
}
func main() {
	//1,只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
	var p1 person

	p1.name = "小红"
	p1.city = "上海"
	p1.age  = 16
	fmt.Printf("p1=%#v\n", p1)

	//2,匿名结构体
	var user struct{name string;age int}
	user.name = "小花"
	user.age  = 17
	fmt.Printf("user=%#v\n", user)

	//3,指针类型结构体
		//通过使用new关键字对结构体进行实例化，得到的是结构体的地址
	var p2 = new(person)
	//对结构体指针直接使用.来访问结构体的成员
	p2.name = "小明"
	p2.city = "上海"
	p2.age  = 18
	fmt.Printf("p2=%#v\n", p2)

	//4,取结构体的地址实例化
		//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	p3 := &person{}
	p3.name =	"张三"
	p3.city	=	"北京"
	p3.age	=	19
	fmt.Printf("p3=%#v\n", p3)
	//p3.name = "张三"其实在底层是(*p3).name = "张三"，这是Go语言帮我们实现的语法糖。

	//5,初始化
	var p4  person
	fmt.Printf("p4=%#v\n", p4)
	//键值对初始化
	p5 := person{
		name: "李四",
		city: "北京",
		age:  20,
	}
	fmt.Printf("p5=%#v\n", p5)

	p6 := &person{
		name: "李四",
		city: "天津",
		age:  21,
	}
	fmt.Printf("p6%#v\n", p6)
	//当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	p7 := &person{
		name: "王子",
		age:  10,
	}
	fmt.Printf("p7%#v\n", p7)
	//6使用值的列表初始
	//初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：
	p8 := &person{
		"小红",
		"北京",
		28,
	}
	fmt.Printf("p8%#v\n", p8)
}
