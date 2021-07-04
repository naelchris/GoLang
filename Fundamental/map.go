package main

import "fmt"

func main() {

	// var numbers map[string]int -> creates a nill map, except the map cannot be added

	// another way to define map
	numbers := make(map[string]int)

	numbers["one"] = 1
	numbers["two"] = 2

	fmt.Println("the first number : ", numbers["one"])

	//init map
	rating := map[string]float32{"C": 5, "Go": 4.5, "C++": 2}

	cRating, ok := rating["C"]

	if ok {
		fmt.Println(cRating)
	} else {
		fmt.Println("there is no things associate with C")
	}

	// map is a reference type. If two maps point to same underlying data, any change will affect both of them.
	m := make(map[string]string)

	m["hello"] = "Hai nael"

	m1 := m
	m1["hello"] = "Hai jessie"

	fmt.Println(m["hello"])

	//when you use var to create map
	var vMap1 map[string]int = make(map[string]int)
	vMap1["hello"] = 2

	fmt.Printf("%d\n", vMap1["hello"])
}
