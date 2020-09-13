package main

import "fmt"
//数组初始化
func main() {

	//我们还可以使用指定索引值的方式来初始化数组，例如
	var testArr = [...]int{1:1,2:5,5:9}


	fmt.Println(testArr)
	fmt.Printf("type of a:%T\n", testArr) //type of a:[6]int
}


