package main

import "fmt"

func main() {
	//array
	coral := [3]string{"black", "blue", "green"}
	fmt.Println(coral)

	//slice
	a := []int{-3, -2, -1, 0, 1, 2, 3}
	fmt.Println(a)

	b := []float64{3.14, 92.34, 111.22, 43.32}
	fmt.Println(b)

	c := []string{"sharks", "golden fish", "cuttlefish"}
	fmt.Println(c)

	//append
	c = append(c, "seahorse")
	fmt.Println(c)
}
