package main

import "net/http"

func main() {
	http.ListenAndServe(":8020", http.FileServer(http.Dir(".")))
}