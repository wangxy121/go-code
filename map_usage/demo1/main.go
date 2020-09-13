package main

import "fmt"

func main() {
	m := make(map[string]int,2)

	m["jack"] = 100
	m["lili"] = 90
	fmt.Println(m)
	fmt.Println(m["jack"])
	fmt.Printf("type of a:%T\n", m)

	v,ok := m["jack"]
	//如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
