package main

import (
	"fmt"
)

func main() {
	var arr [10]int
	arr[0] = 42
	arr[9] = 13

	fmt.Printf("the first element %d\n", arr[0])

	fmt.Printf("the last element %d\n", arr[len(arr)-1])

	doubleArray := [2][4]int{[4]int{1, 23, 4, 4}, [4]int{1, 2, 3, 4}}

	fmt.Printf("%v\n", doubleArray)

	//slicing
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	fmt.Printf("%s\n", ar[2:4])

	//define an array
	var aSlice []byte

	aSlice = ar[:6]

	ar[5] = 'N'

	fmt.Printf("%s\n", aSlice)
	fmt.Printf("%s\n", ar)

	//build in function in slice
	//len -> gets the length of slice
	//cap -> gets the maximum length of slice
	//append
	//copy -> copy elements from one slice to the other, return the number of elements that were copied

}
