package main

import "fmt"

//基于数组通过切片表达式得到切片
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]  // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))

	//为了方便起见，可以省略切片表达式中的任何索引。省略了low则默认为0；省略了high则默认为切片操作数的长度:

	//a[2:]  // 等同于 a[2:len(a)]
	//a[:3]  // 等同于 a[0:3]
	//a[:]   // 等同于 a[0:len(a)]


	//对于数组或字符串，如果0 <= low <= high <= len(a)，则索引合法，否则就会索引越界（out of range）。
}