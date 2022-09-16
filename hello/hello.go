package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hi there!")
	fmt.Println("Please enter your name.")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hi, %s! I'm Go!\n", name)
	fmt.Println("A quote for you: " + quote.Go())
}
