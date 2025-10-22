package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 100 {
		fmt.Println(from, ":", i)
	}
}

// Suppose we have a function call f(s). Here's how we'd call that in the usual way, running it synchronously.

// To invoke this function in a goroutine, use go f(s). This new goroutine execute concurrently with the calling one.

// You can also start a goroutine with an anonymus function all,

// Our two function callas are running asynchronouly in separate goroutines now. Wait for them to finish (for a
// more robust approach, use a WaitGroup)

// When we run this program, we see the output of the blocking call first, then the output of the two goroutines.
// The goroutines output may be interleaved, because goroutines are being run concurrently by the Go runtime.

func main() {
	f("direct")
	go f("goroutine1")
	go f("goroutine2")
	go f("goroutine3")
	go f("goroutine4")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
