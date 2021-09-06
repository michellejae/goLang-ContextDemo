package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/michellejae/go/contextDemo/log"
)

func main() {
	flag.Parse()
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// if we just did below, we could possible 'collide' with another person in the world uses the same key 42 to create IDs
	// so we have to make sure in LOG file to make int42 it's own tyke key or something
	ctx = context.WithValue(ctx, int(42), int64(100))

	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler ended")

	fmt.Printf("value for foo is %v", ctx.Value("foo"))
	select {
	// if receive value of waiting five seconds print hello
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
		// if receive value that the conext.done was called
	case <-ctx.Done():
		// get the context error
		err := ctx.Err()
		// print the context error the server
		log.Println(ctx, err.Error())
		// send the context error back to client as well as the internal server errror
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
