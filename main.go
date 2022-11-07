package main

import (
	"fmt"
	"net/http"
)

func handlerFunc (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"<h1> Welcome to my awesome site! </h1>")
}

func main() {
	fmt.Println("Nameste Go!")
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)// (":3000", handler)
}