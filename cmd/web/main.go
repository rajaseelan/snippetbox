package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a new command line flag with the name addr
	addr := flag.String("addr", ":4000", "HTTP network address")

	// use flag.Parse()to parse the command-line flag
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// create the file server, servers out of the ./ui/static directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// use mux.Handler to register the file server as the handler for
	// all url paths that start with /static/
	// strip "/static" before it reaches the file server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s\n", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
