package main

import "fmt"

func main() {
	a := "hello, 世界"
	for i, c := range a {
		fmt.Println("%d: %s\n", i, string(c))
	}
	fmt.Println("length of hello,世界", len(a))
}
