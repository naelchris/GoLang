package main

import "fmt"

//constant
const PI float32 = 3.1415926
const i = 1000
const MaxThread = 10
const prefix = "astaxie_"

//var
var age int32 = 23

func main() {
	// var <variable name> <Type>
	//short assigment can only be used inside a function and every variable in it has to be used.
	_, b := 32, 32

	var available bool //boolean types
	available = true

	//Numerical Type
	//rune is alias of int32
	//byte is alias of uint8

	// var a int8
	// var b int32

	// c := a + b
	//you cannot do the above codes -> int8 and int32 cannot be assigned together.

	//Float has type float32 and float64

	//--------------------------------------
	//go support complex number as well
	//complex128 -> 64 bit real and imaginary part
	//complex64 -> 32 bit real and 32 bit imaginary part

	var c complex64 = 5 + 51

	//--------------------------------------
	//String

	var frenchHello string
	var emptyString string = ""

	frenchHello = "Bonjour"

	//frenchHello[0] = 'H' //It's impossible to change string values by index. You will get errors when you compile the following code.

	x := []byte(frenchHello)
	x[0] = 'C'
	hello2 := string(x)
	fmt.Printf(hello2)

	// print
	fmt.Println(PI)
	fmt.Println(b)
	fmt.Println(available)
	fmt.Println("Value is : %v", c)

	fmt.Println(frenchHello)
	fmt.Println(emptyString)
}
