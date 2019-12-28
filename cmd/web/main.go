package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/rajaseelan/snippetbox/pkg/models/mysql" // new import

	_ "github.com/go-sql-driver/mysql"
)

// Add snippet field to application struct
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Initialize a mysql.SnippetModel instance and add it to the application
	// dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
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
