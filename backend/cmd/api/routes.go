package main

import (
	"backend/cmd/api/handler"
	"net/http"
)

func routes(h *handler.Handler) {
	http.HandleFunc("/example", handler.Example)
	http.HandleFunc("/example/all", h.ExampleAll)
	http.HandleFunc("/example/delete/", h.ExampleDelete)
	http.HandleFunc("/example/insert", h.ExampleInsert)
}
