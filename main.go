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
	fmt.Fprintf("adding some change")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
