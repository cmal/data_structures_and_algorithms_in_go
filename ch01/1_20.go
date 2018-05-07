package main

import "fmt"

type student struct {
	rollNo int
	name string
}

func main() {
	stud := student{1, "Jonny"}
	fmt.Println(stud)

	fmt.Println("student name: ", stud.name)

	pstud := &stud

	fmt.Println("student name: ", pstud.name)

	fmt.Println(student{rollNo: 2, name: "Ann"})
	fmt.Println(student{name: "Ann", rollNo: 2})
	fmt.Println(student{name: "Alice"})
}
