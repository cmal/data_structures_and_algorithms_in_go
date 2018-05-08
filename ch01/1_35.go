// tower of hanoi

func main() {
	TowersOfHanoi(3)
}

func TOHUtil(num int, from string, to string, temp string) {
	
}

func TowersOfHanoi(num int) {
	TOHUtil(num, "A", "C", "B")
}
