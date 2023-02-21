package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	http.HandleFunc("/home", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequest()
}
