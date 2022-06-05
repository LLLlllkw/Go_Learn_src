package dao

import (
	"fmt"
)

// Simulating a simple biz function

func simpleDAO(input int) string {
	fmt.Println("This is a simple function from dao, the input is ", input)
	return fmt.Sprintln("This is a simple function from dao, the input is ", input)
}
