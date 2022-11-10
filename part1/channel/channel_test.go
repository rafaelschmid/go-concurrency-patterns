//	Kevin Chen (2017)
//	Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

//	Golang channels

package channel

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
CHANNELS
A channel in Go provides a connection between two goroutines, allowing them to communicate.
*/

func TestChannel(t *testing.T) {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // <–c  it will wait for a value to be sent.
	}
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string, c chan<- string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // c <– value, it waits for a receiver to be ready.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

/*
A sender and receiver must both be ready to play their part in the communication.
Otherwise we wait until they are.

Thus channels both communicate and synchronize.

The Go approach
Don't communicate by sharing memory, share memory by communicating.
*/
