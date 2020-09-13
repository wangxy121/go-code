package main

import "fmt"

//使用make()函数构造切片
func main() {
	//make([]T, size, cap)

	a := make([]int,2,10)

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))


	//要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。

	//var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	//s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	//s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil

	//所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。

	//切片的赋值拷贝

	s := []int{1,2,3,4,5}
	s1 := s
	fmt.Println("s value:",s)
	fmt.Println("s1 value:",s1)
	s1[2] = 100
	fmt.Println("s value:",s)
	fmt.Println("s1 value:",s)
	//拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容


}
