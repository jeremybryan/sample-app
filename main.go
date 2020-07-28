package main

import (
	"fmt"
	"log"
	"net/http"
)

const name = "Golang"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Kubernetes ♡ %s!", name)
	})
	fmt.printf("Adding a simple statement here")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
