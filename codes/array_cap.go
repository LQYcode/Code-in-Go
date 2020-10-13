package main

import "fmt"

func main() {
	numbers := make([]int, 4)
	for i := 0; i < cap(numbers); i++ {
		numbers[i] = i
	}
	fmt.Println(numbers)
}
