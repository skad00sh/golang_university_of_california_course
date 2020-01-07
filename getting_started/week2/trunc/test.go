package main

import "fmt"

func main() {
	// This will hold the value that user provide
	var input float32
	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&input)
	// This will convert the number to integer and print
	fmt.Printf("Truncated number is: %d\n", int(input))
}
