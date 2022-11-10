package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func handlerFunc (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w,"To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com \"> support@lenselocked.com </a>.")
}

func home (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprintf(w,"<h1> Welcome to my  awesome site! </h1>")
	if err := homeTemplate.Execute(w,nil); err != nil {
		panic(err)
	}

} 

var homeTemplate *template.Template

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}
	
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/support", handlerFunc)
	
	http.ListenAndServe(":3000",r )
}



/*
Exercise 1.
Add an FAQ page to your application under the path /faq
This does not need to have a lot of content, but make sure it is different from your pages.


*/