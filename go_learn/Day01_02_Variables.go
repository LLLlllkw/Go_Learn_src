package main

import "fmt"

func main() {
	// go 申明变量语法: var <variable> <type>
	// go 要求申明的变量必须被使用，不然报错
	var number int   // go允许申明变量时不赋值，此时number为初始值0
	var number1 = 40 // go 允许申明时不申明类型, 但必须赋值, 赋值时会自动配置变量的类型
	number2 := 0     // 简短申明，直接省略类型和 'var'。 但变量名必须要是一个新的变量名，而不能是已经申明的变量。
	// e.g., number1 := 40 Wrong
	// 这种方式只能用于函数体内，而不能在全局变量中
	// 多变量申明
	var a, b, c = 40, 50, 'a' // 不赋初值时即为默认值
	a, b, c = 1, 2, 3         // 如果之前a,b,c 已经被申明初始化过，那这里的操作就退化成赋值
	var (
		a1 int
		a2 float32
	)
	fmt.Printf("%d,%d,%d", number, number1, number2)
	println(a, b, c)
	println(a1, a2)
}
