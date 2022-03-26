package main

import (
	"fmt"
)

func main() {
	// 1.1 声明一个结构体 go里面没有类的概念，所以用结构体来类比类
	type Student struct {
	}
	type Student1 struct {
		name string
		age  int
	}
	// 1.2 和数组类似， 每一个不同结构体都是一种不同的数据类型，在声明变量时需要声明是哪一种struct的变量
	var b Student
	b = Student{} // 如果 b = Student1{} 是会报错的, b是一个Student类型的变量，而不是Student1
	fmt.Println(b)

	// 1.3 可以声明一个匿名结构体，然后实例化(初始化)的时候不给他传值(后面那个{}就是表示在构建结构体时不传值, 仅使用字段的默认值)
	a := struct {
		name string
		age  int
	}{}
	fmt.Println("a", a)
	// 1.4 结构体数据类型的访问
	fmt.Println("name,", a.name)
	fmt.Println("age,", a.age)

	// 1.5 结构体指针
	// 根据go 的语法糖，可以直接使用结构体指针去修改结构体的属性(目前看到只有struct能这么用)
	ptr := &Student1{name: "zyc", age: 5}
	fmt.Println(ptr)
	ptr.name = "lkw"
	fmt.Println(ptr)

	// 1.6 匿名字段
	// 匿名字段就是不写属性的名字, 只写属性的类型, 那么在我们初始化结构体的时候就必须按顺序写。
	// 在结构体中，匿名字段只能声明不同类型的字段，不能有两个一样类型的字段
	// 它只适用于字段比较少的结构体，以及一些比较简单的场景
	// 并不常用

	// 1.7 如果字段中有结构体的嵌套，那么默认会把被嵌套的结构体中的字段也隐式地加入到这个当前这个结构体

	type Human struct {
		name   string
		age    int
		weight int
	}
	type Student2 struct {
		Human      // 匿名字段  可等效于 <name> Human
		int        // 内置类型作为匿名字段  可等效于 <name> int
		speciality string
	}
	// 在初始化结构体时，如果有匿名字段，可以直接使用类型(type)名作为key。 这也就是为什么匿名字段不能有两个相同的(类型)
	jane := Student2{Human: Human{"Jane", 35, 100}, int: 5, speciality: "Biology"}

	fmt.Println(jane)

	// 1.8 等号比较
	// struct可以用等号比较的条件是struct内部的所有字段都可以比较

}
