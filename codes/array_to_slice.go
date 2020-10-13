package main

import "fmt"

func main() {
	coral := [4]string{"blue coral", "yellow coral", "foliose coral", "black coral"}

	//define slice
	ocean := make([]string, 3)
	fmt.Printf("%q\n", ocean)

	oceans := make([]string, 3, 5)
	fmt.Printf("%q\n", oceans)

	//array[starting index : starting index + length of slice]
	fmt.Println(coral[1:3])
	fmt.Println(coral[:3])
	fmt.Println(coral[1:])

	//array to slice
	coralSlice := coral[:]
	fmt.Println(coralSlice)
	coralSlice = append(coralSlice, "red coral")
	fmt.Println(coralSlice)
	coralSlice = append(coralSlice, "antipathes", "leptopsammia")

	moreCoral := []string{"massive coral", "soft coral"}
	coralSlice = append(coralSlice, moreCoral...)
	fmt.Println(coralSlice)

	coralSlice = append(coralSlice[:3], coralSlice[4:]...)
	fmt.Println(coralSlice)

}
