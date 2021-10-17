package main

import (
	"fmt"

	hello "fmartinez/hellogo"
)

func main() {
	// Get a greeting message and print it.
	message := hello.Greet("Freddy")
	fmt.Println(message)
}
