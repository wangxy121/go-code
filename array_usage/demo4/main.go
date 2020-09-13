package main

import "fmt"
//遍历
func main() {
	var city = [...]string{"上海","北京","深圳"}

	for index,value := range city{
		fmt.Println(index, value)
	}

	for i:=0;i <len(city);i++{
		fmt.Println( city[i])
	}


}
