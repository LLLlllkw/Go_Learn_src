package main

import (
	"fmt"
	"message" // Used to create request struct with PB
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	TimeStamp := time.Now().Unix()
	request := message.OrderRequest{OrderId: "201907300001", TimeStamp: TimeStamp}
	var response message.OrderInfo
	err = client.Call("OrderService.GetOrderInfo", request, &response) // you can use either response or &response as response is already the pointer.
	// the compiler will transform poniter to its value automatically no matter how
	// many pointers are there.
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(response)
}
