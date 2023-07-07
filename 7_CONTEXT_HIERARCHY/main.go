package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// creates a parent context that is cancellable
	parentContext, cancelFunction := context.WithCancel(context.Background())

	childContext, _ := context.WithCancel(parentContext)

	//launch two treatments in another goroutines ; concurrent treatments
	go doSomething(parentContext, "parent")
	go doSomething(childContext, "child")

	//wait 4 seconds
	time.Sleep(time.Second * 4)

	//then cancel all that stuff by using parent context's cancel function
	cancelFunction()

	//then wait another second before exiting the program
	time.Sleep(time.Second)
	fmt.Println("end of program")
}

func doSomething(cancellableContext context.Context, contextName string) {
	//actually it does nothing but wait to be cancelled
	for {
		select {
		case <-cancellableContext.Done():
			fmt.Printf("It's done for context : %s\n", contextName)
			return
		}
	}
}
