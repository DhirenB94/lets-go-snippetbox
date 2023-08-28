package main

import (
	"html/template"
	"net/http"
	"strconv"
)

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
	//read the template file into the template set
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.severError(w, err)
		return
	}
	//execute  to write the template content as the response body
	err = ts.Execute(w, nil)
	if err != nil {
		app.severError(w, err)
		return
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	idFromUrl := r.URL.Query().Get("id")
	if idFromUrl == "" {
		app.notFound(w)
		return
	}
	id, err := strconv.Atoi(idFromUrl)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	app.infoLog.Printf("Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}