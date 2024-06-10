package main

import (
	"net/http"
	"tesgolang/config"
	"tesgolang/controller"
)

func main() {
	config.GetConnection()

	http.HandleFunc("/",  controller.Index)

	http.HandleFunc("/create",controller.Create)

	http.HandleFunc("/delete",controller.Delete)

	fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}