//	Kevin Chen (2017)
//	Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

//  Deterministically quit goroutine with quit channel option in select

package main

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerator(t *testing.T) {
	quit := make(chan bool)
	ch := generator("Hi!", quit)
	for i := 10; i >= 0; i-- {
		fmt.Println(<-ch, i)
	}
	quit <- true
	time.Sleep(time.Second)
}

func generator(msg string, quit chan bool) <-chan string { // returns receive-only channel
	ch := make(chan string)
	go func() { // anonymous goroutine
		for {
			select {
			case ch <- fmt.Sprintf("%s", msg):
				// nothing
			case <-quit:
				fmt.Println("Goroutine done")
				return
			}
		}
	}()
	return ch
}
