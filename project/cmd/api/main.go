package main

import (
	"net/http"
	"project/handler"
)

func main() {
	fs := http.FileServer(http.Dir("static"))

	// Привязка обработчика к URL-префиксу
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	handler.NewHandler()

	http.HandleFunc("/shop/add", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/shop/get", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/shop/delete", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/shop/drop", func(w http.ResponseWriter, r *http.Request) {

	})

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
