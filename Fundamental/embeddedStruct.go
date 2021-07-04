package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	Skills
	int       //build in type as embedded field
	specialty string
}

func main() {
	jane := Student{Human: Human{"Jane", 35, 100}, specialty: "Biology"}
	nael := Student{Human: Human{"Nael", 20, 60}, Skills: []string{"C#", "java", "golang"}, specialty: "CS"}

	fmt.Println("Nael skill : ", nael.Skills)
	fmt.Println("Nael friend : ", jane.name)

	jane.Skills = []string{"neurosurgeon"}
	fmt.Println("Jane Skills : ", jane.Skills[0])
}
