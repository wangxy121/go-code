package main

import "fmt"

type mover interface {
	move()
}
type person struct {
	name string
	age uint8
}

func (p person) move()  {
	fmt.Printf("%s在玩",p.name)
}
func main() {
	var x mover
	p := person{
		name: "小猫",
		age: 9,
	}
	x = p
	x.move()
}
