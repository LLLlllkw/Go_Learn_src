package main

import (
	"fmt"
	"sync"
)

//	1. channel用来给不同的go routine之间共享变量。go 推荐用通信来共享内存， 而不是全局变量 或者其他共享内存的方式
//  2. 通道具有类型， 每一种通道只能通过与其定义时相同的类型，通道的零值是nil

var test chan int          // 创建一个为nil的通道，没有任何用处
var test2 = new(chan int)  // 同上
var test1 = make(chan int) // 类似于map， 需要用make函数，会给对象一个初值。同时返回一个指向该对象的指针。（PS，slice可以理解为返回的是struct）
//  当然我们不用在使用slice，map，channel的时候先去取值在使用，这其实是编译器的一个优化。
//  https://www.cnblogs.com/zhouj-happy/archive/2019/06/02/10962500.html
var test3 = make(map[string]int)

func main() {
	fmt.Println(test)
	fmt.Println(test1)
	fmt.Println(test3)

	//  example of unbuffered channel
	var wg = sync.WaitGroup{}
	// 等待 2个routine完成，如果不等待，因为main也会从通道中取值，当第一次往通道里传值时，无法判断这个值会被
	//  ReadFromChan 还是 main获得，如果被main获得，那main就结束了，导致僵尸协程的出现。
	wg.Add(2)
	go Write2Chan(test1, &wg)
	go ReadFromChan(test1, &wg)
	fmt.Println("TO WAIT MAIN", <-test1)
	wg.Wait()

	// example of close channel
	var IntChan chan int = make(chan int)
	wg.Add(2)
	go SendData(IntChan, &wg)
	go ReadData(IntChan, &wg)
	wg.Wait()

	// example of range
	// 可以看到虽然我在下面写的声明是只写(读)通道，但是在创建和传参时依然使用的是双向通道。因为在创建时，没必要
	// 去创建一个单向通道，只能发或者只能写的通道是没有意义的。
	var Chan1 chan int = make(chan int)
	wg.Add(2)
	go TestRangeWrite(Chan1, &wg)
	go TestRangeRead(Chan1, &wg)
	wg.Wait()
}

func func_test(a chan int) {
	fmt.Printf("%p", a) // 在这里a实际上是个指针，但在使用时也被编译器自动理解成了取值，所以这里也指向原本的chan 对象
}

//  3. 发送，接受数据
//  4. 非缓冲通道
//		非缓冲通道指的是当有一个routine往通道中传入一个对象时，如果暂时没有接收者，则会阻塞当前协程
//		同样， 如果当前需要从通道收到一个值，但是目前没有routine给通道传值的话，也会阻塞当前协程
func ReadFromChan(a chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("开始接收数据")
	res := <-a
	fmt.Println(res)
	fmt.Println("输出完成")

}

func Write2Chan(a chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("输入数据")
	var read int
	fmt.Scanf("%d", &read)
	fmt.Println("传输数据")
	// 当第一次给通道传值时，由于此时有2个协程等待从通道中获取值，究竟给哪个协程是由调度器调度的，无法人工干预。
	a <- read
	a <- read
	fmt.Println("输入完成")
}

//	5. 死锁  fatal error: all goroutines are asleep - deadlock!
//		举个简单的例子，当一个协程在等待一个channel中传来的值，而当前已经没有其他协程会给此通道传值时，就会
//		造成死锁，等待是无效等待，程序会自动结束并抛出错误。
// 			最简单的死锁:
// 				func main(){
// 					chann := make(chan int)
// 					fmt.Println(<-chann)  此时不会再有其他对象给这个channel传值，所以等待无效死锁
// 				}
//		除了channel死锁，在同步中的 sync.WaitGroup 也会死锁，如果我们设定需要等待 3 个routine的结束，而总共
//			就2个routine在运行，最后一个不管等待多久也没有，在这时就会抛出这种异常。
//			例子: 在上面的程序中如果我们给函数传递的是wg对象而不是*wg,就会报这个错误，因为这样的话每个函数中
//			的wg就是原本wg的深拷贝。

//	6. 关闭通道
//		发送方关闭通道即通知接收方发送结束
//		接收方可以通过 v,ok := <-chan 的形式，如果ok为false，表示该通道已关闭，如果从已关闭的通道上取值,
//		取到的是对应通道数据类型的0值
//	6.1 通道关闭的问题
//		当通道关闭后，如果还有routine往通道send，会引起panic
//					 如果关闭一个已经关闭的通道，会引起panic
//		一些rules:
//			不要在消费端关闭channel
//			不要在有多个并行的生产者时关闭channel
//		一些可用的设计模式:
//			当有多个sender时，可以考虑使用计数器的形式，当计数器到0时，才close这个channel，当然这必须要保证计数器
//				的原子性
//			可以用另外一个辅助channel，利用select，当receiver没有通过辅助channel发出stop信号，则继续发送
//	7. range
//		range可以像迭代器一样迭代着从channel中接收数据，range结束的信号是当channel close的时候

func SendData(chann chan int, wg *sync.WaitGroup) {
	for i := 1; i < 10; i++ {
		chann <- i
	}
	close(chann)
	wg.Done()
}

func ReadData(chann chan int, wg *sync.WaitGroup) {
	for {
		if v, ok := <-chann; ok {
			fmt.Println("Value", v)
		} else {
			fmt.Println("channel has closed")
			break
		}
	}
	wg.Done()
}

//	只读通道
func TestRangeRead(chann <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range chann {
		fmt.Println("Value: ", v)
	}
	fmt.Println("channel closed")
}

//	只写通道
func TestRangeWrite(chann chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Println("push Value: ", i)
		chann <- i
	}
	close(chann) // 这里其实不应该用sender去close这个通道，但是为了简便，就凑合一下
}
