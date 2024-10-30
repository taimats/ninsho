package handler

import (
	"backend/pkg/db"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func Example(w http.ResponseWriter, r *http.Request) {
	message := "Hello, world!"
	w.Write([]byte(message))
}

func (h *Handler) ExampleAll(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	/*
		sql, err := db.QueryParse("example/all", nil)
		if err != nil {
			log.Fatalf("Failed parse: %v", err)
		}
	*/

	rows, _ := h.DB.Query("select id, name from example")
	results := []Result{}
	for rows.Next() {
		result := Result{}
		rows.Scan(&result.ID, &result.Name)
		results = append(results, result)
	}

	json.NewEncoder(w).Encode(results)
}

func (h *Handler) ExampleDelete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/example/delete/")
	arg := struct{ ID string }{ID: id}

	sql, err := db.QueryParse("example/delete", arg)
	if err != nil {
		log.Println(err)
	}

	_, err = h.DB.Exec(sql)
	if err != nil {
		log.Println(err)
	}

}
