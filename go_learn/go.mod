module main

require package_test v0.0.0-00010101000000-000000000000

require google.golang.org/protobuf v1.26.0 // indirect

replace package_test => ../package_test

require (
	Inner v0.0.0-00010101000000-000000000000
	ProtoBuf v0.0.0
	github.com/golang/protobuf v1.5.2

)

replace ProtoBuf => ./ProtoBuf

replace Inner => ./ProtoBuf/Inner

go 1.18
