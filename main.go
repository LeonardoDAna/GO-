package main

import (
	"fmt"
)

func main() {
	var age = 14
	fmt.Println("内存地址", &age)
	var cpAge *int = &age
	fmt.Println("hello world", *cpAge)
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}
