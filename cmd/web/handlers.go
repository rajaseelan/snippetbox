package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// change the signature of the home handler fo it is defined as a method
// against *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Use the serverError() helper
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err) // use the serverError() helper
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // use app.notFound helper
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID: %d...", id)
}

// change signature of createSnippet to be defined against
// *application
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	// create some vars to hold dummy data
	title := "O snail"
	content := "O snail\nClimb Moujnt Fuji\nBut slowly slowly!\n\n- Kobayashi Issa"
	expires := "7"

	// Pass the data to SnippetModel.Insert()
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Redirect the use to the relevant page for the snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
