package main

import (
	"fmt"
	"go/types"
)

func main() {
	// 1. Conditional Structure If
	var a int
	fmt.Scanf("%d", &a)
	if a == 0 {
		println("a = 0\n")
	} else {
		println("a != 0\n")
	}
	// 1.2 Another Format
	// 在if 语句中可以创建一个新的变量(e.g., b), 该变量只能在if语句中使用，在外部是访问不到的
	if b := 5; a < b {
		fmt.Printf("a smaller than 5\n")
	} else {
		fmt.Printf("a bigger or equal 5\n")
	}

	// 2. switch
	//     (1). switch 语句不需要使用break 跳出循环
	//     (2). default 不管位置在哪里(即使是第一个匹配，也是最后一个被匹配到的选项)
	//     (3). switch 后面的值可以是任何类型, 即使是一个常量都可以。但case中的类型必须和switch的类型相同, case也可以是一个表达式，
	//    		会计算出对应的值。
	//	   (4). case 中常量不能重复，但是表达式没有check
	switch a {
	case 1:
		fmt.Println("a = 1\n")
	case 2:
		fmt.Println("a = 2\n")
	default:
		fmt.Println("a is out of expectation\n")
	}
	// 2.1
	//    (1) switch 如果想要当某个case执行完后继续执行后面的case，使用fallthrough
	//    (2) switch 语句也和if 一样，可以使用一个局部的变量进行辅助
	switch c := 10; c > 0 {
	case true:
		fmt.Println("c > 0\n")
		fallthrough
	case false:
		fmt.Println("c <= 0\n")
		fallthrough
	default:
		fmt.Println("c is out of expectation\n")
	}
	// 2.2
	//     (1) switch 后面的表达式省略即为true，然后匹配case中为true的对应的case
	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}
	// 2.3 Type switch 先了解 typeswitch可以用来了解一个变量的类型
	var i interface{} // 一个空接口的变量
	switch i.(type) {
	case types.Array:
		fmt.Println("It is an array")
	case int:
		fmt.Println("it is an Int")
	case nil:
		fmt.Println("it is a nil")
	default:
		fmt.Println("ERROR")
	}
}
