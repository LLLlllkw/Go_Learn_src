package service

import (
	"HomeWork03/internal/biz"
	"fmt"
)

func SimpleService() {
	res := biz.SimpleFunction(5)
	fmt.Println("A simple Service, call biz.SimpleFunction, result : ", res)
}
