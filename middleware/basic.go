//go:build basic
// +build basic

package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func BasicMiddleware() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	http.HandleFunc("/", logging(root))

	http.ListenAndServe(":8083", nil)
}

func main() {
	BasicMiddleware()
}
