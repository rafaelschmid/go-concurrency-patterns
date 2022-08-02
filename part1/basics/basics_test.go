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

func TestBoring(t *testing.T) {
	boring("boring!")
}

func TestGoBoring(t *testing.T) {
	go boring("boring!")
}

func TestGoBoringPrint(t *testing.T) {
	go boring("boring!")
	fmt.Println("I'm listening.")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}
