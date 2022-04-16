package main

import "fmt"

type MyStr struct {
	name string
	id   int
}
type MyMap map[string]MyStr

func main() {

	a := make(map[string]MyStr)
	fmt.Printf("%p\n", a)

	ReadMap(&a)

	// var TestMap map[string]func() MyStr = map[string]func() MyStr{}

}

func ReadMap(Map *map[string]MyStr) {
	fmt.Printf("%p\n", Map)
}
