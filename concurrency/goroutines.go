// Concurrency in Go
// Executing functions in parallel and concurrently by using a feature Goroutines
package concurrency

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true //<- sends data(true) through a channel when its done
}

// go - meaing running in parallel as goroutines - running in non blocking way
// dispatch(doesn't wait for function/goroutine to complete) these four goroutines/tasks and we're done/exit from main function = dispatching goroutines is faster than execution/complete these functions
// SOLUTION: channels - path to communicate with goroutines
func ShowGreetExample() {
	dones := make([]chan bool, 4)
	// You can use one and the same channel with multiple Goroutines because this channel is in the end a transmission device. Its capable to be used to receive multiple values.
	// done := make(chan bool)

	dones[0] = make(chan bool)
	go greet("Nice to meet you!", dones[0])
	go greet("How are you?", dones[1])
	go slowGreet("How ... are ... you ...?", dones[2])
	go greet("I hope you're liking the course!", dones[3])

	for _, done := range dones {
		<-done
	}
	// you have to wait for as many values as you have goroutines when you want them to be finished
	// <-done
	// <-done
	// <-done
	// <-done //we waiting for data to come out of the channel and sending it to a void - to go simply means that we're done as soon as we got one value out of this channel.(race condition)
}
