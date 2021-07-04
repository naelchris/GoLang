package main

import "fmt"

type isInt func(int) bool //define the type for function isOdd and isEven

func isOdd(integer int) bool {
	if integer%2 != 0 {
		return true
	}
	return false
}

func isEven(integer int) bool {
	return integer%2 == 0
}

func filter(integer []int, function isInt) []int {
	var result []int
	for _, v := range integer {
		if function(v) {
			result = append(result, v)
		}

	}

	return result
}

func main() {
	var slice = []int{1, 2, 3, 4, 5, 7}

	odd := filter(slice, isOdd)
	even := filter(slice, isEven)

	fmt.Println("slice = ", slice)
	fmt.Println("Odd elements of slice are: ", odd)
	fmt.Println("Even elements of slice are: ", even)
}
