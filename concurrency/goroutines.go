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
	// You can use one and the same channel with multiple Goroutines because this channel is in the end a transmission device. Its capable to be used to receive multiple values.
	done := make(chan bool)

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)
	<-done //we waiting for data to come out of the channel and sending it to a void
}
