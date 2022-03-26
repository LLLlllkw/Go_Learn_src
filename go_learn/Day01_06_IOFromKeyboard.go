package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* IO类似于scanf和 printf
	 1. Println, Printf, Print
		Println 和 Print 类似 打印数值或者字符串
		区别:
			(1) println 打印每一项之间或者空格，类似于Python3 print的 sep 参数，有间隔。 而 print没有
			(2) println 每次输出完会换行, 类似于Python3 print的 end='\n'。 而 print 没有
		Printf 用于格式化输出，和C 类似 同时Printf也不会自动换行
			比较特殊的一些占位符
				%T 直接打印类型
				%v 原样输出
				%c 打印字符
				%p 打印地址
				%t bool
	*/
	a := 65
	fmt.Printf("type of a is %T, value of a is %d, address of a is %p \n", a, a, &a)
	fmt.Printf("%x \n", a)

	/*
		 2. Scanf Scan 和 Scanln
				都是输入，多个输入参数时在terminal用空格分割不同的输入变量
				Scanf 是格式化输入，在占位符字符串里的其他字符也要输入，不然就报错。即输入和format的形式不一致，就会报错
					同时遇到回车会立即结束
				Scan 和 Scanln区别
					scan 遇到回车不会结束输入 会把回车也看做空白分割符，直到全部需要输入的变量都输入为止
					scanln 遇到回车即结束，如果有还没输入的变量，就使用默认值
	*/
	var (
		i int
		j int
		k int
	)
	args, err := fmt.Scanln(&i, &j, &k)
	fmt.Println(args)
	fmt.Println(err)
	/*
		 3. Bufferio
		    需要先创建reader 对象
			然后根据需要读取的对象类型去使用方法
	*/
	reader := bufio.NewReader(os.Stdin)
	stringRead, _ := reader.ReadString('\n')
	fmt.Println(stringRead)

}
