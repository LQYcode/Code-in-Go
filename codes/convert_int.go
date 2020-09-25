package main

import "fmt"

func main() {
	var index int8 = 15
	var bigIndex int32

	bigIndex = int32(index)
	//convert int & know %T
	fmt.Println(bigIndex)
	fmt.Printf("index data type:  %T\n", index)
	fmt.Printf("bigIndex data type:  %T\n", bigIndex)

	//convert big to little
	var big int64 = 64
	var little int8
	little = int8(big)
	fmt.Println(little)

	//maximum int
	var big1 int64 = 129
	var little1 = int8(big1)

	fmt.Println(little1)

}
