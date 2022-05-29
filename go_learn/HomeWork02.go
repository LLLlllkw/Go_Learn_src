package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/pkg/errors"
)

var ServerOut = make(chan struct{}) // a conn channel to inform other go routine a server crash occured
var OsOut = make(chan struct{})     // a conn channel to inform other go routine a os exit occured

/*
This demo does not really work as server.ListenAndServe method will block and the only method for that is server.Shutdown
Which requires a context
The part of context is still learning
The demo uses channel to send signal between goroutines
*/
func main() {
	var g = errgroup.Group{}
	mux := http.NewServeMux()
	mux.HandleFunc("/err/", ServerErr)
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	g.Go(func() error {
		fmt.Println("start server goroutine")
		err := server.ListenAndServe()
		if err != nil {
			return err
		}
		select {
		case <-OsOut:
			fmt.Println("meet os out")
			return errors.Errorf("get os signal")
			// case <-ServerOut:
			// 	fmt.Println("meet server out")
			// 	return errors.Errorf("get server signal")
		}
	})
	// os signal
	g.Go(func() error {
		fmt.Println("start os signal goroutine")
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case sig := <-quit:
			OsOut <- struct{}{}
			return errors.Errorf("get os signal: %v", sig)
		case <-ServerOut:
			fmt.Println("receive Server exit signal")
			return errors.Errorf("get server signal")
		}
	})
	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}

func ServerErr(w http.ResponseWriter, req *http.Request) {
	fmt.Println("simulate an error!\n")
	ServerOut <- struct{}{}
	runtime.Goexit() // it will only terminate the handler go routine, but not for the routine where listen_and_serve locates
}
