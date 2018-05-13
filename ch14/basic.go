package main

import (
	"fmt"
	_ "strings"
	"reflect"
)

func main() {
	s := "hello World!"
	r := []rune(s)
	r[0] = 'H'
	s2 := string(r)
	fmt.Println(s2)
	s3 := "你好"

	fmt.Println(reflect.TypeOf(s2[0]), reflect.TypeOf(s), reflect.TypeOf(s2[:4]), reflect.TypeOf(s3[0]))
}
