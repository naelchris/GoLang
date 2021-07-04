package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

//inheritance
// If an anonymous field has methods, then the struct that contains the field will have all the methods from it as well.

type Student struct {
	Human  //anonymous field
	school string
}

type Employee struct {
	Human   //anonymous field
	company string
}

//method in Human Object, Employee and student will inherit this method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)

}

//override method SayHi for Employee
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}

	sam.SayHi()
	mark.SayHi()
}
