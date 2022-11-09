package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFunc (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w,"To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com \"> support@lenselocked.com </a>.")
}



func home (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w,"<h1> Welcome to my  awesome site! </h1>")

}


// Ques 1
func faqPage (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1> This is faq page :) </h1>")
}




func homeNotFoundHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w,"<h1> Caustion : Under Construction site..! </h1>")

}

 

func main() {
	
	r := mux.NewRouter()
	var h http.Handler = http.HandlerFunc(homeNotFoundHandler)

	r.HandleFunc("/", home)
	r.HandleFunc("/support", handlerFunc)
	r.HandleFunc("/faq", faqPage)
	
	r.NotFoundHandler = h
	http.ListenAndServe(":3000",r )
}

/*
Exercise 1.
Add an FAQ page to your application under the path /faq
This does not need to have a lot of content, but make sure it is different from your pages.


*/