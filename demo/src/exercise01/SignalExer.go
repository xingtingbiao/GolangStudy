package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
信号的练习
*/
func main() {
	//testSignal()
	testExit()
}

func testSignal() {
	// Go signal notification works by sending `os.Signal`
	// values on a channel. We'll create a channel to
	// receive these notifications (we'll also make one to
	// notify us when the program can exit).
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		sig := <-signals
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

func testExit() {
	// `defer`s will _not_ be run when using `os.Exit`, so
	// this `fmt.Println` will never be called.
	// 当主动调用os.Exit(1)退出时, defer将不被执行
	defer fmt.Println("!")

	// Exit with status 3.
	os.Exit(3)
}
