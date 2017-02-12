package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/events", handler)
	http.ListenAndServe(":5050", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {

	writer.WriteHeader(http.StatusOK)
}