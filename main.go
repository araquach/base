package main

import (
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"html/template"
	"net/http"
)

var (
	tplHome *template.Template
	tplInfo *template.Template
	tplRegister *template.Template
)

func init() {
	tplHome = template.Must(template.ParseFiles("views/layouts/main.gohtml", "views/pages/home.gohtml"))
	tplInfo = template.Must(template.ParseFiles("views/layouts/main.gohtml", "views/pages/info.gohtml"))
	tplRegister = template.Must(template.ParseFiles("views/layouts/main.gohtml", "views/pages/register.gohtml"))

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/info", info)
	r.HandleFunc("/register", register)

	// The path "/" matches everything not matched by some other path.
	http.Handle("/", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplHome.Execute(w, nil); err != nil {
		panic(err)
	}
}

func info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplInfo.Execute(w, nil); err != nil {
		panic(err)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplRegister.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	appengine.Main() // Starts the server to receive requests
}