package inner_package

import "fmt"

func AddTwoNumbers(a, b int) (res int) {
	fmt.Println("I am now in a new package")
	return a + b
}
