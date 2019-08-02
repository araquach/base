package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func routes() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/info", info).Methods("GET")
	r.HandleFunc("/freelance", freelance).Methods("GET")
	r.HandleFunc("/apprentice", apprentice).Methods("GET")
	r.HandleFunc("/registera", registera).Methods("GET")
	r.HandleFunc("/registerf", registerf).Methods("GET")
	r.HandleFunc("/create", create).Methods("POST")
	r.HandleFunc("/success", success).Methods("GET")

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

	return r
}


