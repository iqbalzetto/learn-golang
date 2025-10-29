package main

import (
	"fmt"
	"time"
)

//We can use channels to synchronize execution across goroutines. Here's an example of
//using a blocking receive to wait for a goroutine to finish. When waiting for multiple
//goroutines to finish, you may prefer to use a Waitgroup

// this function we'll run in a goroutine. The done channel will be used to notify
// another goroutine that this function's work is done
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	//send a value to notify that we're done
	done <- true
}

func main() {
	done := make(chan bool, 1)
	//start a worker goroutine, giving it the channel to notify on
	go worker(done)

	//block until we recieve a notification from the worker on the channel
	<-done

	//if you remove the <-done line from this program, the program will exit before
	//the worker finished its work, or in some cases even before it started
}
