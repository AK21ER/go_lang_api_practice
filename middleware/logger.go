package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc { // here http.handlerfunc  is a type that represents a function that can handle HTTP requests. It takes an http.ResponseWriter and an *http.Request as parameters and returns nothing. By using this type, we can ensure that any function we pass to the Logger middleware will have the correct signature to handle HTTP requests.
	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		// BEFORE request
		fmt.Println("➡️ Request:", r.Method, r.URL.Path)

		// call actual handler
		next(w, r)

		// AFTER request
		fmt.Println("⬅️ Completed in:", time.Since(start))
	}
}