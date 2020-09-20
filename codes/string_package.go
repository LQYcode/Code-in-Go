package main

import (
	"fmt"
	"strings"
)

func main() {
	//ss := "Sammy Shark"
	//fmt.Println(strings.ToUpper(ss))
	//fmt.Println(strings.ToLower(ss))

	//string search
	//fmt.Println(strings.HasPrefix(ss,"Sammy"))
	//fmt.Println(strings.HasSuffix(ss,"Shark"))
	//
	//fmt.Println(strings.Contains(ss,"Sh"))
	//fmt.Println(strings.Contains(ss,"sh"))
	//
	//fmt.Println(strings.Count(ss,"S"))
	//fmt.Println(strings.Count(ss,"s"))

	//c := [] string {"sharks","golden fish","cuttlefish"}
	//fmt.Println(strings.Join(c,","))

	balloon := "Summy has a balloon"
	s := strings.Split(balloon, " ")
	fmt.Println(s)
	fmt.Printf("%q", s)

	fmt.Println(strings.ReplaceAll(balloon, "has", "had"))

}
