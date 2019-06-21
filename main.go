package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var (
	tplHome *template.Template
	tplInfo *template.Template
	tplRegister *template.Template
)

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

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/public/images/favicon.ico")
}

func main() {
	port := os.Getenv("PORT")
	// port := "5050"
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	tplHome = template.Must(template.ParseFiles("views/layouts/main.gohtml", "views/pages/home.gohtml"))
	tplInfo = template.Must(template.ParseFiles("views/layouts/main.gohtml", "views/pages/info.gohtml"))
	tplRegister = template.Must(template.ParseFiles("views/layouts/main.gohtml", "views/pages/register.gohtml"))

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/info", info)
	r.HandleFunc("/register", register)

	// Styles
	assetHandler := http.FileServer(http.Dir("./dist/"))
	assetHandler = http.StripPrefix("/dist/", assetHandler)
	r.PathPrefix("/dist/").Handler(assetHandler)

	// JS
	jsHandler := http.FileServer(http.Dir("./dist/"))
	jsHandler = http.StripPrefix("/dist/", jsHandler)
	r.PathPrefix("/dist/").Handler(jsHandler)

	//Images
	imageHandler := http.FileServer(http.Dir("./dist/images/"))
	r.PathPrefix("/dist/images/").Handler(http.StripPrefix("/dist/images/", imageHandler))


	http.HandleFunc("/favicon.ico", faviconHandler)

	// The path "/" matches everything not matched by some other path.
	http.Handle("/", r)
	http.ListenAndServe(":" + port, r)
}