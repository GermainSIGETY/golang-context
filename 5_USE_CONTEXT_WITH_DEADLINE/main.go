package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Create a context with a deadline feature : 'at that very moment, everything will be done Buddy'
	// almost same stuff than context with a timeout ; instead of duration it's datetime
	// it returns :
	// a context with a way to check whether 'it's done or not done' (see below)
	// a cancel function too (underscored here because we won't use it : see directory 2_USE_CONTEXT_WITH_CANCEL)
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*4))

	//launch a treatment in another go routine ; concurrent treatment
	go doSomething(ctx)

	//wait 10 seconds
	time.Sleep(time.Second * 10)
	fmt.Println("end of program")
}

func doSomething(contextWithTimeout context.Context) {
	for {
		select {
		// A context expose a 'Done()' function. This function returns a channel
		// when an element is read from this channel 'it's the signal that it's done Buddy';
		// you can stop and trash everything related to your current treatment.
		case <-contextWithTimeout.Done():
			fmt.Println("It's Done Buddy")
			return
		default:
			fmt.Println("I will wait 1 second")
			time.Sleep(1 * time.Second)
		}
	}
}
