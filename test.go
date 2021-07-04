package main

import "fmt"

type Employee struct {
	Human
	department string
	status     string
}

type Human struct {
	name string
	age  string
	food string
}

func (h *Human) eat(food string) {
	h.food = food
}

func main() {
	Nael := Employee{Human{name: "nael", age: "20"}, "cs", "single"}

	Nael.eat("burger")

	fmt.Printf(Nael.food)
}
