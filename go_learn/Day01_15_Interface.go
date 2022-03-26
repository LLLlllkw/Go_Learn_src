package main

import "fmt"

func main() {
	// interface可以被任何对象实现，但这种实现是隐式的，不需要像java 一样用implements
	// 任何类型都实现了空接口 interface{}
	// 在结构体中如果有结构体嵌套，被嵌套的结构体实现了某个接口，那外部的结构体也自动继承了这个接口的实现

	// interface 变量
	// 如果声明了一个interface 变量，那么这个变量可以接收任何实现了这个接口的对象
	// ATTENTION: 这里有一个注意点。
	//     如果在实现接口时的接收者是指针, 那该接口变量也只能接收该对象的指针
	//     如果在实现接口时的接收者是对象本身, 那该接口变量也只能接收该对象本身, 而不能是指针
	//     也就是说, 接口变量能接收的对象种类(一般就指针和对象本身两种) 必须和实现接口的方法的接收者的种类相同。
	var zyc PersonAbility    // 定义了一个interface 变量
	zyc = &Girl{name: "ZYC"} // 如果没有取地址, 则错误
	fmt.Println(zyc)

}

// 1.1 定义一个接口，类似指定了一个protocol, 实现该接口，就是实现接口中定义的所有方法
type PersonAbility interface {
	Sayhi() // 函数名 (<入参>) <出参>
	Talk()
}

type Girl struct {
	name string
}

func (person *Girl) Sayhi() {
	fmt.Println("Hi I am a girl")
}

func (person *Girl) Talk() {
	fmt.Println("Hi I am a girl and I can Talk")
}
