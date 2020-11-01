package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Golang")
}

func main() {
	http.HandleFunc("/", index)
	PORT := ":3000"
	log.Println("Application running on port", PORT[1:])
	log.Fatal(http.ListenAndServe(PORT, nil))
}
