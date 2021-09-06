package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	// context.Background will NEVER be canclled -- will be used one ALL EXAMPLES
	ctx := context.Background()

	// this is a child of context.Background so i think you have to have it called first
	// WILL BE USED WITH TWO AND THREE
	//ctx, cancel := context.WithCancel(ctx)

	// FOUR
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	defer cancel()

	// ONE
	// if we do not do anything, the hello should print. however, if we start typing somehing, our cancel will
	// trigger and we will cancel the context (ie the program )
	// go func() {
	// 	s := bufio.NewScanner(os.Stdin)
	// 	s.Scan()
	// 	cancel()
	// }()

	// TWO
	// please cancel context after one second ... sleepAndTalk is never called if so
	// go func() {
	// 	time.Sleep(time.Second)
	// 	cancel()
	// }()

	// THREE
	// same as above
	//time.AfterFunc(time.Second, cancel)

	// FOUR

	sleepAndTalk(ctx, 5*time.Second, "hello")
}

// funcion evaluates what value it recieves first ... if it gets the value sent via the time package it will
// print our message. if it receives the value from the context it will print the error
// take use of channels from the packages to send the value
func sleepAndTalk(ctx context.Context, d time.Duration, s string) {
	select {
	case <-time.After(d):
		fmt.Println(s)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
