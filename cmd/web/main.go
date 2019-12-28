package main

import (
	// new import
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // new import
)

// Define an application struct to hold application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	// define a new cmd-line flag for MysQL DSN string
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// keep the main() function clean
	// code for creating connection pool in separate openDB() function
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// close db before main() exits
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(), // call the new app.routes() method
	}

	infoLog.Printf("Starting server on %s\n", *addr)
	// err already declared, so reuse
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

// openDB() function wraps the sql.Open() and returns a sql.DB connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
