package main

/*
For errors in DAO(DAL), the error is from lower function call (e.g., from mysql module), which only has a simple santinel error.
We need to add context to find out which query results in the error for debug.
*/

import (
	BasicErrors "errors"
	"fmt"

	"github.com/pkg/errors"
)

var (
	mock_db         = map[string]string{"a": "1", "b": "2", "c": "3"}
	ErrNoRows error = BasicErrors.New("ErrNoRows") // a santinel error without stack record
)

type QueryReq struct {
	req string
}

// The basic Query simulating Mysql package
func Query(req QueryReq) (string, error) {
	if _, ok := mock_db[req.req]; !ok {
		return "0", ErrNoRows // For basic module, returns a santinel error
	}
	return mock_db[req.req], nil
}

// Simulating DAO(DAL)
// DAO(DAL) need some wrap info to record stack info and failed input for debug
func QueryCaller(req QueryReq) (string, error) {
	res, err := Query(req)
	if err != nil {
		return "0", errors.Wrapf(err, "DAL Query Failed, Failed request %+v", req)
	}
	return res, nil
}

// Simulating a call from application
// The DAL will return error with additional context, instead of handling it
// You need to decide how to solve it in application layer
func main() {
	var req QueryReq = QueryReq{req: "1"}
	res, err := QueryCaller(req)
	if err != nil {
		fmt.Printf("Receive Err. Trace route: \n%+v\n", err)
		if errors.Is(err, ErrNoRows) {
			fmt.Println("It is an ErrNoRows")
		}
		return
	}
	fmt.Println("Query res:", res)
}
