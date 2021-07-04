package main

import "fmt"

type Human struct {
	name  string
	age   string
	phone string
}

type Employee struct {
	Human
	specialty string
	phone     string
}

func main() {
	Nael := Employee{Human{"Nael", "20", "089992"}, "Backend", "33-22"}
	fmt.Println(Nael)
}
