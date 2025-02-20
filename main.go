package main

import (
	"log"
	"net/http"
)

func main() {
	//fmt.Println("Hello World")

	http.HandleFunc("GET /ping", hanleping)

	log.fatal(http.ListenAndServe(":8080", nil))
}

func hanleping(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
}
