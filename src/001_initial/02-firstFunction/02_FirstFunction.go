package main

import (
	"fmt"
)

func main() {

	var a = 10
	var b = 20
	fmt.Println(add(a,b))
}

func add(a int, b int) int {
	return a + b
}