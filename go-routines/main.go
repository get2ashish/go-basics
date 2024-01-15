package main

import (
	"fmt"
	"time"
)

func greet(phrase string, isDoneChannel chan bool) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, isDoneChannel chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	isDoneChannel <- true
}

func main() {
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	<-done
	<-done
	<-done
	<-done
	//go will wait till this is done
}
