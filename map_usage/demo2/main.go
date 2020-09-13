package main

import "fmt"

func main() {
	m := make(map[string]int,2)

	m["jack"] = 100
	m["lili"] = 90
	delete(m,"lili")
	for index,value := range m{
		fmt.Println(index,value)
	}


}
