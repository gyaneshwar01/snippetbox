package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// create a file server that serves files out of the "./ui/static" directory
	// path given to http.Dir function is relative to project directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// For matching paths, we strip the "/static" prefix from the request before it reaches to fileServer
	// If we don't strip the prefix, it looks for ./ui/static/static/... which doesnot exist
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// mux.HandleFunc() is a syntactic sugar that transforms the function into http.HandlerFunc and registers in single step
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
