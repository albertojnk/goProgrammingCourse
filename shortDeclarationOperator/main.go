package main

import "fmt"

var z = 21

func main() {
	x := 42
	y := "James Bond"
	fmt.Println(x, y)
	x = 50
	fmt.Println(x)
	fmt.Println(z)
}

func foo() {
	fmt.Println(z)
}
