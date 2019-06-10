package main

import (
	"google.golang.org/appengine"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template


func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)

	// The path "/" matches everything not matched by some other path.
	http.Handle("/", r)
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles(
		"views/home.gohtml")
	if err != nil {
		panic(err)
	}

	appengine.Main() // Starts the server to receive requests
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}