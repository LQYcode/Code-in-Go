package main

import "fmt"

func main() {
	//a := `say "hello" to Go!`
	var a string = `say "hello" to Go`
	fmt.Println(a)

	b := `say "hello" to Go! \n`
	fmt.Println(b)

	c := `this string is on multiple lines
		within a single back quote on both side.
		`
	fmt.Println(c)

	d := "say \"hello \"  to Go!"
	fmt.Println(d)

}
