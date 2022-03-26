package main

import (
	"fmt"
	"math"
)

func main() {
	//func funcName(parametername type1, parametername type2) (output1 type1, output2 type2) {
	//这里是处理逻辑代码
	//返回多个值
	//return value1, value2
	//}
	print2int(1, 2)
	PrintAllInt(1, 2, 3, 4, 5, 5, 6, 7, 7, 1, 2)
	nu := 9.0
	fmt.Println("原本值所在的内存地址", &nu)
	// 传值
	// 这里的函数没有给名字, 但也不会报错，因为有一个变量去接收了这个函数对象
	getSquareRoot := func(x float64) float64 {
		x = math.Sqrt(x)
		return x
	}
	// 传地址
	getSquareRootAdd := func(x *float64) float64 {
		fmt.Println("指针所在的内存地址", &x) // 该地址和原本nu所在的地址不同
		fmt.Println("指针所指向的内存地址", x) // 该地址和原本nu所在的地址相同
		*x = math.Sqrt(*x)
		return *x
	}
	//  !! Important !! go 其实都是值传递，没有引用传递，即使是指针，也是把这个指针复制一份新的，从而能访问到原始的变量值

	/* 使用函数 */
	fmt.Println(getSquareRoot(nu))
	fmt.Println(nu)
	fmt.Println(getSquareRootAdd(&nu))
	fmt.Println(nu)

	// 函数对于返回值也可以给一个变量名，这样就不需要用return语句返回这个值了, 但仍然需要显式的给出return语句
	getSquareRootNew := func(x float64) (xAfterSquareRoot float64) {
		xAfterSquareRoot = math.Sqrt(x)
		return
	}
	fmt.Println(getSquareRootNew(9.0))

	fmt.Println("---------Decorator----------")
	// 函数本身也是一种数据类型，所以可以作为参数传递给其他函数
	// 函数数据类型和数组类似, 数组的大小是数组类型的一部分。 函数这里的话函数的入参和返回值类型必须完全一致才是同一个类型，否则就是
	// 不同类型
	// 这里搞了一个闭包, 闭包可以在外层函数退出时保留外层函数的局部变量, 如果这个变量被内层函数使用的话，
	// 外层函数的变量对于内层函数来说就相当于全局变量，是作用域高一级的变量。
	// 下面这个例子在外层函数中传递了getSquareRoot 这个函数，然后在内层函数中调用了这个函数
	closure := func(old_inner func(number1 float64) float64) func(number1 float64) (NewResult float64) {
		fmt.Println("received inner function is ", old_inner) // 可以看出这个变量名是个指针, 指针内存中存放的值和getSquareRoot一致，但是指针本身的内存地址变化了
		return func(number1 float64) (NewResult float64) {
			fmt.Println("call old inner function")
			NewResult = old_inner(9.0)
			fmt.Println("the following is new ")
			return
		}
	}
	fmt.Println(getSquareRoot)
	fmt.Println(closure(getSquareRoot)(9.3))

	// defer
	// 用来延迟某些操作，被defer的语句只有在最后整个函数执行完,在return之前的时候才会执行
	// 如果有多个defer, 采用栈的形式，后进先出
	ShowDefer()
	// defer 语句中的参数是在执行defer的时候的值，而不是等到真正执行这个语句时候的值
	// defer 可以理解成保存现场的作用，当执行到defer语句时，将函数相关的信息保存起来存储到一个stack里，到最后再pop出来
	n := 5
	defer ShowDeferVar(n) // 打印出的是5 而不是10
	n = 10
	// defer函数：
	//  当外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行。
	//  当执行外围函数中的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回。
	//  当外围函数中的代码引发运行恐慌时，只有其中所有的延迟函数都执行完毕后，该运行时恐慌才会真正被扩展至调用函数。

}

func print2int(number1 int, number2 int) {
	fmt.Println("number1", number1)
	fmt.Println("number2", number2)
	return
}

// 可变参数 func myfunc(arg ... <type>)  是生成一个slice来存放所有可变参数
func PrintAllInt(arg ...int) {

	for index, number := range arg {
		fmt.Println("index:", index, "number:", number)
	}

}
func ShowDefer() {
	a := 1
	b := 2
	defer fmt.Println("b: ", b)
	fmt.Println("a: ", a)
}

func ShowDeferVar(a int) {
	fmt.Println("a:", a)
}
