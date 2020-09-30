package main

import "fmt"

func main() {
	var x int64 = 57
	var y float64 = float64(x)

	fmt.Printf("%.2f\n", y)

	var f float64 = 390.8
	var i int = int(f)

	fmt.Printf("f = %.2f\n", f)
	fmt.Printf("i = %d\n", i)

	var d float64 = 50.0
	var e float64 = 12

	g := d / e
	fmt.Println(g)

}
