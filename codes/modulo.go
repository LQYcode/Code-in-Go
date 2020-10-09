package main

import (
	"fmt"
	"math"
)

func main() {
	o := 85
	p := 15
	fmt.Println(o % p)

	q := 36.0
	r := 8.2
	s := math.Mod(q, r)
	fmt.Println(s)
}
