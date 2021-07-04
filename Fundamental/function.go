package main

import "fmt"

func main() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	for key, element := range employee {
		fmt.Println("key : ", key, "element : ", element)
	}

	for _, v := range employee {
		fmt.Println("map val: ", v)
	}

	fmt.Printf("%d\n", max(2, 3))

	var x, y = sumAndProduct(2, 3)
	fmt.Printf("%d %d\n", x, y)

	myfunc(1, 2, 3, 4, 5)

	var num []int = []int{12, 3, 4, 55, 3}
	myfunc(num...)
} //end of main

func funcName(input1 int, input2 string) string {
	return input2
}

func max(n int, m int) int {
	if n > m {
		return n
	}
	return m
}

//multivalue return -> return results of A + B and A * B
func sumAndProduct(a int, b int) (add int, multiplied int) {
	add = a + b
	multiplied = a * b
	return
}

//varadic function parameter as argument
func myfunc(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for v, n := range nums {
		fmt.Println(v)
		total += n
	}
	fmt.Println(total)
}

//pointer function -> input memeory address of the variable (&a)
func edit(a *int) int {
	*a = *a + 1
	return *a
}

//defer wait until the function return then reverse it
