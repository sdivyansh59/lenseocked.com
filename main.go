package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handlerFunc (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, r.URL.Path)
	
	if(r.URL.Path == "/") {
		fmt.Fprintf(w,"<h1> Welcome to my  awesome site! </h1>")
	} else if (r.URL.Path == "/support") {
		fmt.Fprintf(w,"To get in touch, please send an email to <a href=\"mailto:support@lenslocked.com \"> support@lenselocked.com </a>.")
	}else{
	
		
		// w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w,"<h1> We did not find what you are looking for :)</h1> <br> please send us a mail.")
		
	}
	
	
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}



func main() {

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	

	// http.HandleFunc("/", handlerFunc)
	// http.ListenAndServe(":3000", nil)
	http.ListenAndServe(":3000",router )


}

	/* (":3000", handler)
	 nil means: we want default serve mux */


	/* fresh for dynamic reload server 
	 go get github.com/pilu/fresh  */

// for routing we are using github.com/julienschmidt/httprouter