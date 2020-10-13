package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	error := errors.New("barnacles")
	fmt.Println("Sammy says:", error)

	newErr := fmt.Errorf("error occurred at: %v", time.Now())
	fmt.Println("An error happened:", newErr)
}
