package main

import (
	"fmt"
	"time"
)

func main() {
	doSomething("Mensagem do e-mail")
	// filterOldEmails()
}

//	To use the concurrency (do something in parallel), we use the "go" keyword

func doSomething(message string) {
	go func() { //	This function block will be executed separetedly
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("E-mail recebido: %v\n", message)
	}() // -> All "go" functions must be a call
	fmt.Printf("E-mail enviado: %s\n", message)
}

//	CHANNELS
//	Channels are a typed, thread-safe queue. They allow different goroutines to communicate with each other
//	To create a channel, we do: "ch := make(chan T)"
//	And to send data to that channel: ch <- 69
//	And to receive data from a channel: v := <-ch

func filterOldEmails() {
	//	Do not run. The syntax said above is correct, but this function has a big mistake when using channels and go routines:
	//	It deadlocks. A deadlock is when all goroutines in a program are sleeping, so they have no data to pass.
	isOldChan := make(chan bool)

	isOldChan <- true

	isOld := <-isOldChan

	fmt.Print(isOld)

}

//	Empty structs are often called "tokens" in Go programs. In this context, a token is a unary value.
//	In other words, we don't care what value is passed through the channel. We cara when and if is passed.
//	We can block and wait until it something is sent on a channel using the following syntax:
//		<-ch || ...chan struct{}
//	This will block until it pops a single item off the channel, then continue, discarding the item

//	BUFFERED CHANNELS
//	Channels can optionally be buffered.
//	It's possible to provide the buffer length as the second argument in the make() function
//		ch := make(chan int, 100)
//	Sending on a buffer channel only blocks when the buffer is full
//	Receiving blocks when the buffer is empty

//	CLOSING A CHANNEL
//	Channels can be explicitly closed by a sender.
//		ch := make(chan int)
//		do some stuff
//		close(ch)
//	To check if a channel is closed, it's similar to when accessing a map:
//		v, ok := <-ch
//		ok is false if the channel is empty and closed
//	Do not send on a closed channel. This will cause a panic. A panic on the main goroutine will cause the entire program to crash, whereas a panic on another goroutine will cause that goroutine to crash.
