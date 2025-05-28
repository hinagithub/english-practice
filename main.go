package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Cloud Run!")
}

func main() {
	http.HandleFunc("/", handler)
	port := "8080"
	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, nil)
}
