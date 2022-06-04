package main

import (
	"fmt"
	"net/http"
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World</h1>")
}

func main() {
	http.HandleFunc("/", healthcheck);
	http.ListenAndServe(":8081", nil)
}
