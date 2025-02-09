package main

import (
	"fmt"

	"learning-golang/learning"
)

// testing the main function
func main() {
	fmt.Println("Hello, this is the main function!")

	// call the function from learning package
	learning.DataType()

	// call the function from learning package
	learning.Variable()

	// call the function from learning package
	learning.Convertion()
}
