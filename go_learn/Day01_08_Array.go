package main

import "fmt"
import "reflect"

func main() {
	// 1. 数组声明
	//    (1). var <variable name> [length] [type]
	var a [5]int // 数组没有给其赋值，即为初始值, 对int类型来说是 0
	fmt.Println("array a: ", a)
	var b = [5]int{1, 2, 3, 4, 5} // 数组被赋了初值
	fmt.Println("array b: ", b)
	var c = [...]rune{'a', 'b', 'c', 'd'} // 不设定数组大小，则会根据后面的真实数组大小来设置
	fmt.Println("array c: ", c)
	d := "hello? 你好吗"
	fmt.Println(d[0], d[8])
	fmt.Printf("%c,%c\n", d[0], d[8]) // d[0] 正确输出，d[8]错误，因为字符串的底层是byte数组。一个中文在utf占3bytes
	e := []rune(d)
	fmt.Println("format", reflect.TypeOf(e))
	// e[0] 正确输出，e[8]正确输出, 将字符串(byte数组)转成了rune(int)类型数组，支持utf8的所以字符。这里面编译器有一个自动转换的逻辑
	fmt.Printf("%c,%c\n", e[0], e[8])
	// 用 [...] 来表示让编译器自动调整数组长度
	var f = [...]int{'a'}
	fmt.Println(f)
	// 在数组赋值过程中可以使用<index: value>的形式对数组进行部分初始化
	var g = [9]int{0: 5, 3: 8} // index = 0 为5, index = 3 为8
	fmt.Println(g)
	//总结:
	//go中数组声明如果要赋初始值，那必须使用等号赋值语句，而不能只使用类似 'var a [5]int{1,2,3}'的语句

}
