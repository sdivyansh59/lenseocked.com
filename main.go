package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFunc (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprintf(w, r.URL.Path)
		fmt.Fprintf(w,"To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com \"> support@lenselocked.com </a>.")
}



func home (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w,"<h1> Welcome to my  awesome site! </h1>")

}



func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/support", handlerFunc)

	// router.GET("/hello/:name", Hello)
	// http.HandleFunc("/", handlerFunc)
	// http.ListenAndServe(":3000", nil)

	http.ListenAndServe(":3000",r )


}

	/* (":3000", handler)
	 nil means: we want default serve mux */


	/* fresh for dynamic reload server 
	 go get github.com/pilu/fresh  */

// for routing we are using github.com/julienschmidt/httprouter