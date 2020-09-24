package main

import "fmt"

var g1 = "global"

func printLocalVar() {
	l := "local"
	fmt.Println(l)
	fmt.Println(g1)
}

func main() {
	printLocalVar()
	fmt.Println(g1)
}
