package function

// perbandingan dengan for looping biasa
func FunctionRecursiveLoop(value int) int {

	hasil := 1
	for i := value; i > 0; i-- {
		hasil *= i
	}
	return hasil

}

// FunctionRecursive is a function to show how to create a recursive function in golang
func FunctionRecursive(value int) int {

	if value == 1 {
		return 1
	} else {
		return value * FunctionRecursive(value-1)
	}

}
