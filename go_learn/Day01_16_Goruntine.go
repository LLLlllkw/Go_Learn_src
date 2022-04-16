package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	//  1. goroutine 是go中实现并发的方式，相当于是协程 coroutine, 由于协程对于内核属于透明不感知的, 可以从代码层面
	//         直接调度
	//  2. 为了方便直接调度, go 提供了runtime 包来提供调度的能力
	//  3. 临界资源: 如果一个资源同时被多个routine使用，会导致panic的产生，这时候需要使用锁的机制(大部分其他语言)，或者
	//		使用channel用通信的方式去共享资源
	//		3.1 WaitGroup
	//			用来等待一系列的Goroutine执行完毕, 相当于一个计数器。具有 Add, Wait, Done 三个方法
	//			在一开始设置一个等待, Add() 接受一个int参数表示需要等待多少个goroutine结束
	//			当在某个线程调用wait后，就会阻塞住当前线程。
	//			每次执行完毕一个Goroutine，就会调用Done 方法，使得计数器-1，直到计数器为0，则停止等待，继续执行当前coroutine
	//	如果直接执行这两个线程，是不会有输出的，因为main这个协程会继续往下，而如果main结束了，其他的子协程也无法继续(类似于僵尸)
	//	所以必须要使用groupwait，让main等待这两个routine实行完，才会继续
	wg.Add(4) // 设置计数器需要完成的goroutine数量
	go Rountine3()
	go Rountine4()
	//	4. 锁 虽然go一般不推荐使用锁，但是还是了解一下
	//		4.1 互斥锁
	//			即同一时间只能由一个线程获得这把锁，其他在竞争中失败的线程会等待
	var mutex sync.Mutex // 如果直接 := 的话，需要自己给这个结构体赋初值。 或者用 mutex := new(sync.Mutex)
	go Rountine5(&mutex)
	go Rountine6(&mutex)
	//		4.2 读写锁
	//			读写就类似于数据库了，他的逻辑比较复杂，分为读和写。
	//			因为如果大家同一时间都在读，是没问题的，而只要有一个人在写，就会有问题。所以在有一个人写的时候，大家都不能读，在有一个人读的时候，就不能有人写
	//				读锁不能阻塞读锁
	// 				读锁需要阻塞写锁，直到所有读锁都释放
	// 				写锁需要阻塞读锁，直到所有写锁都释放
	// 				写锁需要阻塞写锁
	//			rwmutex := new(sync.RWMutex)
	//			至于是使用写锁还是读锁，在lock时使用
	//				rwMutex.Rlock 和  rwMutex.WLock()

	wg.Wait()

}

func UseGoRuntime() {
	fmt.Println(runtime.NumCPU())       // 返回Cpu核数
	fmt.Println(runtime.GOMAXPROCS(1))  // 设置Go运行时最多能使用的CPU个数
	fmt.Println(runtime.NumGoroutine()) // 返回正在执行和排队的任务总数
	fmt.Println(runtime.GOOS)           // 操作系统

}

func Rountine1() {
	runtime.Gosched() // 让当前线程让出cpu，供其他线程使用
	fmt.Println("I am in Rountine 1")

}

func Rountine2() {
	runtime.Goexit() // 退出当前线程，但是defer语句会继续执行
	fmt.Println("I am in Rountine 2")
	defer fmt.Println("This is defer")
}

func Rountine3() {
	for i := 1; i < 10; i++ {
		fmt.Println("I am in Rountine 3")
	}
	wg.Done()
}

func Rountine4() {
	for i := 1; i < 10; i++ {
		fmt.Println("I am in Rountine 4")
	}
	wg.Done()
}
func Rountine5(mutex *sync.Mutex) { //  注意一定要指针，具体原因自己想一下
	mutex.Lock() // 上锁，如果没有拿到锁就会等待，这里由于mutex是结构体，虽然传的是指针，但还是可以直接调用其方法
	for i := 1; i < 10; i++ {
		fmt.Println("I am in Rountine 5")
	}
	mutex.Unlock() // 结束后释放
}

func Rountine6(mutex *sync.Mutex) {
	mutex.Lock()
	for i := 1; i < 10; i++ {
		fmt.Println("I am in Rountine 6")
	}
	mutex.Unlock()
}
