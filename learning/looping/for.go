package looping

import (
	"fmt"
)

// ForLooping is a function to show the looping in Golang
func ForLooping() {

	count := 10
	for count >= 1 {
		fmt.Println(count)
		count--
	}

	// bisa juga dijadikan satu lint atau short
	for count := 10; count >= 1; count-- {
		fmt.Println(count)
	}

}
