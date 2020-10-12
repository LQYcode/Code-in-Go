package main

import "fmt"

func main() {
	permissions := map[int]string{1: "read", 2: "write", 4: "create", 8: "delete", 16: "add"}
	delete(permissions, 16)
	fmt.Println(permissions)
	permissions = map[int]string{}
	fmt.Println(permissions)
}
