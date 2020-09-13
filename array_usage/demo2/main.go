package main

import "fmt"
//数组初始化
func main() {
	//一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，例如：

	var testArr [3]int
	var numArr = [...]int{1,2} //使用指定的初始值完成初始化
	var cityArr = [...]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化

	fmt.Println(testArr)                      //[0 0 0]
	fmt.Println(numArr)                       //[1 2 0]
	fmt.Printf("type of numArray:%T\n", numArr)   //type of numArray:[2]int
	fmt.Println(cityArr)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArr) //type of cityArray:[3]string

}
