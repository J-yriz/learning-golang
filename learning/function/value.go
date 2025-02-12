package function

import (
	"fmt"
)

func SayGoodBye(firstName string) string {
	return fmt.Sprintf("Good Bye %s", firstName)
}

// function value is a function to show how to create a function with value in golang
// sebuah function yang disimpan ke dalam variable
func FunctionValue(firstName string) {

	sayBye := SayGoodBye
	fmt.Println(sayBye(firstName))

}
