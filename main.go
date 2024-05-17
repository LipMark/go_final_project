package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("TODO_PORT")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("fine")
}
