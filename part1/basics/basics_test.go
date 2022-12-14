package basics

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

/*
Calling boring function synchronously
Program waits boring function
*/
func TestBoring(t *testing.T) {
	boring("boring!")
}

/*
Calling boring function asynchronously
Program doesnt wait boring function
*/
func TestGoBoring(t *testing.T) {
	go boring("boring!")
}

/*
Calling boring function asynchronously
Program doesnt wait, but has a 5 seconds sleep
*/
func TestGoBoringPrint(t *testing.T) {
	go boring("boring!")
	fmt.Println("I'm listening.")
	time.Sleep(5 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}

/*
GOROUTINE
Goroutine is an independently executing function
It's not a thread.
There might be only one thread in a program with thousands of goroutines.
Instead, goroutines are multiplexed dynamically onto threads as needed to keep all the goroutines running.

The main function couldn't see the output from the other goroutine.
We need communication: channels
*/
