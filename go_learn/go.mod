module main

require package_test v0.0.0-00010101000000-000000000000

require (
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect

)

replace package_test => ../package_test

require (
	Inner v0.0.0-00010101000000-000000000000
	ProtoBuf v0.0.0
	github.com/golang/protobuf v1.5.2
	github.com/pkg/errors v0.9.1
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	google.golang.org/grpc v1.46.2

)

replace ProtoBuf => ./ProtoBuf

replace Inner => ./ProtoBuf/Inner

go 1.18
