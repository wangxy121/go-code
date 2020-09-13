package main

import "fmt"
//遍历
func main()  {
	s := []int{1,2,3,4,5}

	for i := 0; i<len(s);i++  {
		fmt.Println(i,s[i])
	}

	fmt.Println("===")

	for index,value := range s  {
		fmt.Println(index, value)
	}

}
