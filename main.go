package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, Chat app")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}
