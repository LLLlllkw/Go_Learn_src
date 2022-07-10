package main

import (
	"fmt"
)

type test struct {
	a int
	b int
}

func (t test) func_test() {
	if t.c == 1 {
		fmt.Println("a = 1")
	}
	return
}

func main() {
	a := test{a: 1, b: 2}
	a.func_test()
}
