package main

import "fmt"

func addTwoNumbers(x, y int) int {
	sum := x + y
	return sum
}

func multiplyTwoNumbers(x, y int) int {
	product := x * y
	return product
}

func main() {
	m := multiplyTwoNumbers(5, 9)
	fmt.Println(m)
}
