package main

import "fmt"

func foo() {

	//last in will be first out
	defer fmt.Println("Done!")                //last
	defer fmt.Println("are we really done ?") //first

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Doing some stuff")
}

func main() {
	foo()
}
