// Concurrency in Go
// Executing functions in parallel and concurrently by using a feature Goroutines
package concurrency

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true //<- sends data(true) to a channel
}

// go - meaing running in parallel as goroutines - running in non blocking way
// dispatch(doesn't wait for function/goroutine to complete) these four goroutines/tasks and we're done/exit from main function = dispatching goroutines is faster than execution/complete these functions
// SOLUTION: channels - path to communicate with goroutines
func ShowGreetExample() {
	// go greet("Nice to meet you!")
	// go greet("How are you?")
	done := make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	// go greet("I hope you're liking the course!")
	<-done //we waiting for data to come out of the channel and sending it to a void
}
