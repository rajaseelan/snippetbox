package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// use log.New to create a logger for writing informational messages
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create logger for errors
	// use stderr & Lshortfile for the file name and line number
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

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

	// initialize a new http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Write messages using the two new loggers instead of the standard logger
	infoLog.Printf("Starting server on %s\n", *addr)
	// Call the ListenAndServe() method on our new http.Server struct
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
