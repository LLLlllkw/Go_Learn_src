package main

import (
	"fmt"
	"math"
	"net"
	"net/http"
	rpc "net/rpc"
	"os"
)

type MathUtil struct {
}

func (*MathUtil) CalculateCircleArea(radius float32, area *float32) error {
	*area = radius * radius * math.Pi
	return nil
}

func main() {
	// example of RPC without protobuf
	// protobuf (marshall and unmarshall is not a must in rpc)
	mathUtil := new(MathUtil) //初始化指针数据类型

	//2、调用net/rpc包的功能将服务对象进行注册
	err := rpc.Register(mathUtil)
	if err != nil {
		panic(err.Error())
	}

	//3、通过该函数把mathUtil中提供的服务注册到HTTP协议上，方便调用者可以利用http的方式进行数据传递
	rpc.HandleHTTP()

	//4、在特定的端口进行监听
	listen, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(os.Stdout, "%s", "start connection")
	http.Serve(listen, nil) // 最好不要使用go routine，使用的话还需要用waitgroup去上锁，不如直接让他阻塞
	// 使用方法: 开一个shell, go run server.go； 再重新开一个shell, Go run client.go
}
