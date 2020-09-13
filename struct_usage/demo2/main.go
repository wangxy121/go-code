package main

import "fmt"

type person struct {
	name string
	city string
	age   int8
}

func main() {
	//构造函数
	p9 := newPerson("张三", "北京", 20)
	fmt.Printf("%#v\n", p9) //

	p1 := newPerson("jack","北京",20)
	fmt.Println(p1.age)
	p1.setAge(10)
	fmt.Println(p1.age)


	p2 := newPerson("lili","上海",18)
	fmt.Println(p2.age)
	p2.setAge2(25)
	fmt.Println(p2.age)

	p2.Dream()

}

//构造函数
func newPerson(name string,city string,age int8) *person  {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

//方法 和接收者
//Dream Person做梦的方法
func (p person) Dream() {
	fmt.Printf("%s的梦想是好好学习！\n", p.name)
}
//使用值接收者
func (p person) setAge(newAge int8) {
	p.age = newAge
}

//使用指针接收者
func (p *person) setAge2(newAge int8) {
	p.age = newAge
}
