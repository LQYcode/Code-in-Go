package main

import "fmt"

func main() {

	a := "i love you"
	b := []byte(a)
	c := string(b)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
