package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	// port from outside
	// .env file => loader

	// check data
	// handlers === > controllers
	// academ
	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("fine")
}
