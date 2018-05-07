package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Apple"] = 40
	m["Banana"] = 30
	m["Mango"] = 50

	for k,v := range m {
		fmt.Print("[", k, " -> ", v, "]")
	}

	fmt.Println("Apple price: ", m["Apple"])
	delete(m, "Apple")

	v, ok := m["Apple"]
	fmt.Println("Apple price: ", v, " Present: ", ok)

	v2, ok2 := m["Banana"]
	fmt.Println("Banana price: ", v2, " Present: ", ok2)
}
