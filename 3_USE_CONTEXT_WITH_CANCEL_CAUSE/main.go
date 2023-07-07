package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Create a context with a cancellation feature
	// it returns :
	//- a context with a way to check whether it's done or not' (see below)
	//- a cancel function : 'call it and everything will be cancelled'
	cancellableContext, cancelFunction := context.WithCancelCause(context.Background())

	//launch a treatment in another go routine ; concurrent treatment
	go doSomething(cancellableContext)

	//wait 4 seconds
	time.Sleep(time.Second * 4)

	//then cancel all that stuff, and precise a cause
	cancelFunction(fmt.Errorf("crap, it has failed in main function"))

	//then I wait another second before exiting the program
	time.Sleep(time.Second)
	fmt.Println("end of program")
}

func doSomething(cancellableContext context.Context) {
	for {
		select {
		// A context exposes a 'Done()' function, that returns a channel :
		// when an element is read from this channel,
		// 'it's the signal that it's done Buddy';
		// you can stop and trash everything related to your current treatment.
		case <-cancellableContext.Done():
			cause := context.Cause(cancellableContext)
			fmt.Printf("It's Done Buddy, beause of : %s\n", cause)
			return
		default:
			fmt.Println("I will wait 1 second")
			time.Sleep(1 * time.Second)
		}
	}
}
