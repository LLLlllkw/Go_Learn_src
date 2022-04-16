package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//	buffered channel
	//	对于非缓冲通道，当有一个goroutine push 进通道时，会以阻塞的方式等待 另一个 goroutine将其pop
	//	而对于缓冲通道，即我们在传输之前预设了这个通道的大小，如果当前channel中数据长度小于预设长度，
	//	sender可以非阻塞的继续执行下面的代码。如果当前channel数据长度大于预设，则继续以阻塞的方式。
	//	然而对于receiver，如果当前channel中没有数据，那还是会被阻塞。
	Chan := make(chan int, 100) // 创建一个buffered channel
	ChanUBuf := make(chan int)
	wg := &sync.WaitGroup{}
	wg1 := &sync.WaitGroup{}
	// example of buffered
	// 结果可以看出， send routine一次性把所有数据都push进去了，而不会阻塞，等接收方接一个发一个
	wg.Add(2)
	go TestBufChanSend(Chan, wg)
	go TestBufChanRec(Chan, wg)
	wg.Wait()
	// example of unbuffered
	// 从结果可以看出来，他的rec 和 send 是间断的，也就证明是先发一个再收一个。
	wg1.Add(2)
	go TestUBufChanSend(ChanUBuf, wg1)
	go TestUBufChanRec(ChanUBuf, wg1)
	wg1.Wait()

}

func TestBufChanSend(chann chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println("send ", i)
		chann <- i
	}
	close(chann)
}
func TestBufChanRec(chann chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1)
	for {
		v, ok := <-chann
		if ok {
			fmt.Println("rec ", v)
		} else {
			fmt.Println("closed")
			break
		}
	}
}

func TestUBufChanSend(chann chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println("send ", i)
		chann <- i
	}
	close(chann)
}
func TestUBufChanRec(chann chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// time.Sleep(1)
	for {
		v, ok := <-chann
		if ok {
			fmt.Println("rec ", v)
		} else {
			fmt.Println("closed")
			break
		}
	}
}
