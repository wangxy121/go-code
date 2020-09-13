package main

import "fmt"
//数组初始化
func main() {
	//数组定义  var 数组变量名 [元素数量]T
	//比如：var a [5]int， 数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。 [5]int和[10]int是不同的类型。
	//数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1，访问越界（下标在合法范围之外），则触发访问越界，会panic。

	//初始化
	var testArr [3]int//数组会初始化为int类型的零值
	var numArr = [3]int{1,2} //使用指定的初始值完成初始化
	var cityArr = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化

	fmt.Println(testArr)                      //[0 0 0]
	fmt.Println(numArr)                       //[1 2 0]
	fmt.Println(cityArr)                      //[北京 上海 深圳]

}
