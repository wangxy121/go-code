package main

import "fmt"

func main() {
	//x.(T) 类型断言
	//x：表示类型为interface{}的变量
	//T：表示断言x可能是的类型。
	var x interface{}
	x = "Hello 沙河"

	fmt.Println(x.(string))
}
