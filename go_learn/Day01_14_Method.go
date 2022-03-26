package main

import "fmt"

// 结构体的嵌套实际上就是一种继承的体现
type Person struct {
	gender string
}

type Student struct {
	Person
	name string
	age  int
}

func main() {
	// method 方法是基于结构体(类)的。就是在普通的函数之前给一个接收者
	zyc := Student{Person: Person{"male"}, name: "ZYC", age: 5}
	zyc.Sayhi()

	// 接收者为指针和值的区别
	zyc.ChangName()
	fmt.Println(zyc.name)
	zyc.ChangNameAddr()
	fmt.Println(zyc.name)
}

// 基于Student结构体(类)， 给其增添了一个新方法，于是所有Student结构体(类)创建出来的对象(实例化)都可以使用
func (student Student) Sayhi() {
	fmt.Println("Hi I am", student.name)
}

// 由于继承，在父结构体(类)中的方法在外层的结构体(类比于子类)中也可以使用
func (person Person) Say() {
	fmt.Println("I am a Person")
}

// 方法的接收者是结构体和结构体指针的区别
//     方法的接收者这一块的逻辑其实就是在调用这个方法时先把已经初始化的结构体对象(类比于instance)传给这个方法
//     而由于go只有值传递没有引用传递，所有如果接收者是struct值对象, 那就是复制这个值对象一份
//     如果传的是结构体指针，那就复制这个结构体指针
//         当然，在调用的时候你不需要去显示的用指针或者struct对象去调用这个方法
//         e.g., p = struct{}{}; p.setval() 具体是传指针还是是由你声明的这个方法的接收者声明决定的，编译器自动进行了转换
//         所以，在声明方法时，不能声明两个同名方法，即使他们的接收者是结构体和结构体指针
//        		func (student Student) Sayhi () {
//					fmt.Println("Hi I am", student.name)
//					}
//				func (student *Student) Sayhi () {
//					fmt.Println("Hi I am", student.name)
// 					}
//			这样写是会报错的，因为其接收者还是同一个结构体(类)
// 			具体使用结构体指针还是结构体只是一个结构体方法的两种不同写法

func (student Student) ChangName() {
	student.name = "LKW" // 这种是不能修改原本结构体对象中的name字段的，因为传值后是一个原本结构体对象的深拷贝
}

func (student *Student) ChangNameAddr() {
	student.name = "LKW" // 这种是可以修改原本结构体对象中的name字段的，因为传过去一个复制的指针，然后根据语法糖自动寻找到对应的结构体对象
}

//  方法可以在继承中被重写，在调用时就近查找

//  任意类型都可以有方法, 不一定是结构体类型
//  但是必须要求只能对本地类型(本包内)的类型定义方法，不能给别的包的类型定义方法
//  所以如果我们想给int类型添加方法，必须新添加一个自定义类型然后继承自int, 这也体现了封装性
