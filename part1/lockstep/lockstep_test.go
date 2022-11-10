//	Kevin Chen (2017)
//	Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

//	Golang channels as handles on a service (working in lockstep)

package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGenerator(t *testing.T) {
	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 10; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string { // returns receive-only channel
	ch := make(chan string)
	go func() { // anonymous goroutine
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}

/*
These programs make Joe and Ann count in lockstep (alternating between one and the other).

We can instead use a fan-in function to let whosoever is ready talk.
*/
