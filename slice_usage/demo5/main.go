package main

import "fmt"

//append 添加
func main() {
	var s []int
	s = append(s, 1)        // [1]
	s = append(s, 2, 3, 4)  // [1 2 3 4]
	s2 := []int{5, 6, 7}
	s = append(s, s2...)    // [1 2 3 4 5 6 7]

	fmt.Println(s)
	fmt.Println(s2)

}
