package main

import "fmt"

//定义一个函数 它的返回值 也是函数

func Test() func(string) {

	return func(name string){
		fmt.Println("hello",name)
	}


}
func main() {
	f := Test()

	f("www")
}
