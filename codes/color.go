package main

import (
    "fmt"
)

const favColor string = "blue"

func main(){
    var guess string
    //create an input loop
    for{
        //ask the user guess my favorite color
        fmt.Println("Guess my favorite color")

        //try to read from input
        if _, err := fmt.Scanln(&guess); err != nil {
            fmt.Printf("%s\n",err)
            return
        }

        //did they get the correct color?
        if favColor == guess {
            //got it
            fmt.Printf(" %q is correct",favColor)
            return
        }
        //wrong,guess again
        fmt.Printf(" %q is not!")
    }
}