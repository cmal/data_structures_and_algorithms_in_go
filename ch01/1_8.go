package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("something else")
	}
}
 
yuzhao-linux 福 ~/gits/data_structures_and_algorithms_in_go/ch01 
7385 ◯ : [?2004hls
ls[?2004l
1_10.go  1_14.go  1_8.go  1_9.go
 
yuzhao-linux 福 ~/gits/data_structures_and_algorithms_in_go/ch01 
7386 ◯ : [?2004hgo run 1_14.go
go run 1_14.go[?2004l
# command-line-arguments
./1_14.go:9:2: undefined: fmt
 
yuzhao-linux 福 ~/gits/data_structures_and_algorithms_in_go/ch01 
7387 ◯ : [?2004hgo run 1_14.go
go run 1_14.go[?2004l
Sum is ::  55
 
yuzhao-linux 福 ~/gits/data_structures_and_algorithms_in_go/ch01 
7387 ◯ : [?2004hgo run 1_15.go
go run 1_15.go[?2004l
zsh: correct '1_15.go' to '1_14.go' [nyae]? n
n
stat 1_15.go: no such file or directory
 
yuzhao-linux 福 ~/gits/data_structures_and_algorithms_in_go/ch01 
7388 ◯ : [?2004h[?2004l
 
yuzhao-linux 福 ~/gits/data_structures_and_algorithms_in_go/ch01 
7388 ◯ : [?2004hgo run 1_18.go
go run 1_18.go[?2004l
Value stored at variable var is  10
Value stored at variable var is  10
The address of variable var is  0xc4200120c8
The address of variable var is  0xc4200120c8
 
yuzhao-linux 福 ~/gits/data_structures_and_algorithms_in_go/ch01 
7389 ◯ : [?2004hgo run 1_19.go
go run 1_19.go[?2004l
Value of i before increment is :  10
Value of i after increment is :  11
 
yuzhao-linux 福 ~/gits/data_structures_and_algorithms_in_go/ch01 
7390 ◯ : [?2004hgo run 1_20.go
go run 1_20.go[?2004l
{1 Jonny}
student name:  Jonny
student name:  Jonny
{2 Ann}
{2 Ann}
{0 Alice}
 
yuzhao-linux 福 ~/gits/data_structures_and_algorithms_in_go/ch01 
7391 ◯ : [?2004h