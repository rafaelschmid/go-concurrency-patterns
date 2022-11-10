//	Kevin Chen (2017)
//	Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

//	Select is a control structure for concurrency (why channels/goroutines are built in; not library)
//  Based off of Dijkstra's guarded commands... providing an idiomatic way for concurrent processes to
//  pass in data without programmer having to worry about 'steps'

package main

import (
	"fmt"
	"testing"
	"time"
)

/*
PATTERN - SELECT
It's like a switch, but each case is a communication:
- All channels are evaluated.
- Selection blocks until one communication can proceed.
- If multiple can proceed, select chooses pseudo-randomly.
- A default clause, if present, executes immediately if no channel is ready.
*/

func TestGenerator(t *testing.T) {
	ch := fanIn(generator("Hello"), generator("Bye"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

// fanIn is itself a generator
func fanIn(ch1, ch2 <-chan string) <-chan string { // receives two read-only channels
	new_ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				new_ch <- s
			case s := <-ch2:
				new_ch <- s
			default:
				fmt.Println("No one is ready")
				time.Sleep(time.Second)
			}
		}
	}()
	return new_ch
}

func generator(msg string) <-chan string { // returns receive-only channel
	ch := make(chan string)
	go func() { // anonymous goroutine
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(2 * time.Second)
		}
	}()
	return ch
}
