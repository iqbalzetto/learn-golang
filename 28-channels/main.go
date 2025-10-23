package main

import "fmt"

func main() {
	messages := make(chan string)

	// Channels are the pipes that connect concurrent goroutines. you can send values into channels
	// from one goroutine and receive those values into another goroutien.

	// Create a new channel with make(chan val-type). Channels are typed byt the values they convery

	// The <- channel syntax receives a value from the channel. Here we'll receive the "ping" message
	// we sent above and print it out.

	// When we run the program the "ping" message is successfully passed from one goroutine to another via our channel

	// By default sends and receives block until both the sender and receiver are ready. This property
	// allowed us to wait at the end of our program for the "ping" message without having to use any other
	// synchronization

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

}
