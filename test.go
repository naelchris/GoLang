package main

import "fmt"

type Employee struct {
	Human
	department string
	status     string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Human struct {
	name string
	age  string
	food string
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	//Guzzle(beerStein string)
}

//method
func (h Human) SayHi() {
	fmt.Printf("Hi I am %s I am %s years old\n", h.name, h.age)
}

//method
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//method
func (h *Human) eat(food string) {
	h.food = food
}

func main() {
	Nael := Employee{Human{name: "nael", age: "20"}, "cs", "single"}

	//define interface
	var i Men
	i = Nael

	i.Sing("you you")
	Nael.eat("Pizza") //i don't have the method for eat

	fmt.Printf(Nael.food)
}
