package main

import (
	"Inner"
	"ProtoBuf"
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
)

func main() {
	msg_test := &ProtoBuf.Person{
		Name: proto.String("Davie"),
		Age:  proto.Int(18),
		From: proto.String("China"),
	}

	//序列化
	msgDataEncoding, err := proto.Marshal(msg_test)
	if err != nil {
		panic(err.Error())
		return
	}

	msgEntity := ProtoBuf.Person{}
	err = proto.Unmarshal(msgDataEncoding, &msgEntity)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	fmt.Printf("姓名：%s\n\n", msgEntity.GetName())
	fmt.Printf("年龄：%d\n\n", msgEntity.GetAge())
	fmt.Printf("国籍：%s\n\n", msgEntity.GetFrom())
	Inner.PrintInner()
}
