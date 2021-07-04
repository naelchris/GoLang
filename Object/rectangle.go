package main

type Rectangle struct {
	width, height float64
}

// func area(r Rectangle) float64 {
// 	return r.width * r.height
// }

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}

	//by function
	// area(r1)
	// area(r2)

	//by method
	r1.Area()
	r2.Area()

}
