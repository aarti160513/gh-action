package main

import (
	"fmt"
	"os"
)

func main() {
	// Print "Hello, World!"
	fmt.Println("Hello, World!")

	// Check if command-line arguments are passed
	if len(os.Args) > 1 {
		fmt.Println("Command-line arguments:")
		// Loop through the arguments and print them
		for i, arg := range os.Args[1:] {
			fmt.Printf("Argument %d: %s\n", i+1, arg)
		}
	} else {
		fmt.Println("No command-line arguments provided.")
	}
}
