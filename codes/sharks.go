package main

import (
    "fmt"
)

func main(){
    // Define sharks variable as a slice of
    sharks := []string{"hammerhead","great","red"}

    // For loop that iterates over sharks
    for _, shark := range sharks{
        fmt.Println(shark)
    }

}