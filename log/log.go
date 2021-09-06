package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

// this basically gives a special ID for each request so if a request decides to kill early we can know which one it is
// why we have to use context i have no idea,

// key is a type that only the log pakcage can use to help avoid collision with possibly another perosn in the world creating key = 42
type key int

const requestIDKey = key(42)

// get a vaule from the context to identify what we are printing?
func Println(ctx context.Context, msg string) {

	id, ok := ctx.Value(requestIDKey).(int64)

	// if value is not there or wrong type, we will print error
	if !ok {
		log.Println("could not find request ID in context")
		return
	}
	// or it will print our id and the message
	log.Printf("[%d] %s", id, msg)
}

// receiving context, addin ga value, and sending it back
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)
		f(w, r.WithContext(ctx))
	}
}
