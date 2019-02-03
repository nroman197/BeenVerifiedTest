// author: Nazareth Roman
package main

import (
	"net/http"
)

func main() {
	// get mux to handle functions
	mux := getMux()
	// listening port
	http.ListenAndServe("localhost:8000", mux)
}
