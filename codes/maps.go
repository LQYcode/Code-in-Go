package main

import (
	"fmt"
	"sort"
)

func main() {
	sammy := map[string]string{"name": "Sammy", "color": "blue", "animal": "shark", "location": "ocean"}

	fmt.Println(sammy["color"])

	var keys []string

	for key := range sammy {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	fmt.Printf("%q\n", keys)

	items := make([]string, len(sammy))

	var i int
	for _, v := range sammy {
		items[i] = v
		i++
	}
	fmt.Printf("%q\n", items)

	counts := map[string]string{}
	count, ok := counts["sammy"]
	if ok {
		fmt.Printf("Sammy has a count of %d\n", count)
	} else {
		fmt.Println("Sammy was not found")
	}

}
