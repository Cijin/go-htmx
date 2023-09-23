package main

import (
	"log"
	"net/http"

	"github.com/cijin/go-htmx/handlers"
)

const port = ":3000"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/contacts", handlers.RenderContactsPage)
	mux.HandleFunc("/contacts/", handlers.HandleContact)

	mux.HandleFunc("/", handlers.RenderHome)

	log.Fatal(http.ListenAndServe(port, mux))
}
