package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := strconv.Itoa(12)

	fmt.Printf("%q\n", a)

	//user := "Sammy"
	//lines := 50
	//fmt.Println("Congratulations, " + user + ".you have just wrote "+strconv.Itoa(lines)+" lines of code")

	fmt.Println(fmt.Sprint(421.034))
	f := 5524.342
	fmt.Println(fmt.Sprint(f))

}
