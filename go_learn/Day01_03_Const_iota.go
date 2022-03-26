package main

func main() {
	// const <variable> <type> = value
	// 常量在使用过程中不能被修改, 因为不能被修改，所以在初始化时一定要给他赋值
	const c1 int = 'A' // 指定类型
	const c2 = 'a'     // 不指定类型
	const c3 = 111     // const 常量可以仅申明但不使用
	//c1 = 'B'  Cannot assign to c1
	println(c1, c2)
	// 使用常量块申明时，如果第一个常量赋值了但后面的常量没有赋值，后面的常量会自动拥有和第一个常量同样的值
	const (
		e1 = 200
		e2
		e3
	)
	/* ================================= */
	// iota是一个特殊的常量，可以被编译器修改
	// 它是一个枚举值，可以用来很方便的定义枚举
	// 逻辑: iota在const关键字出现时将被重置为0(const内部的第一行之前)，
	// const申明块中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
	// 即: iota是在const什么语句块中用来做自增操作，const语句块每多加一行iota就++1
	const (
		a1 = iota  // a1 = 0
		a2  // a2 = 1
		a3  // a3 = 2
		a4  // a4 = 3
		a5 = iota  // a5 = 4
		a6  // a6 = 5
	)
	const b1 = iota // b1 = 0
	const b2 = 5  // b2 = 0, 因为const关键词出现,同时const又必须赋值，所以iota终止

	const(
		d1 = iota  // 0 iota = 0
		d2  // 1 iota = 1
			// 全空行无效 iota = 1
		d3  //  d3 = 2 iota = 2
		d4 = 200  // iota = 3 但已经有申明具体的值，所以不使用iota的值 d4 = 200
		d5 = iota // 恢复iota的值，iota = 4, d5 = 4
	)


	println(a1,a2,a3,a4,a5,a6)
	println(b1,b2)
	println(d1,d2,d3,d4,d5)
}
