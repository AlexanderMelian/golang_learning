package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Bienvenido a GNUNO")
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "Solicitud recibida - GNUNO")
	} else {
		fmt.Fprintf(w, "404")
	}
	fmt.Println("An User joined  " + r.Method)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":6666", nil))
}

func main() {
	handleRequests()
}
