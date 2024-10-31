package handler

import (
	"backend/pkg/db"
	"backend/pkg/util"
	"encoding/json"
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

	sql, err := db.QueryParse("example/all", nil)
	if err != nil {
		util.WriteErrorJson(w, err)
		return
	}

	rows, err := h.DB.Query(sql)
	if err != nil {
		util.WriteErrorJson(w, err)
		return
	}
	results := []Result{}
	for rows.Next() {
		result := Result{}
		rows.Scan(&result.ID, &result.Name)
		results = append(results, result)
	}

	util.WriteJson(w, http.StatusOK, results, nil)
}

func (h *Handler) ExampleDelete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/example/delete/")
	arg := struct{ ID string }{ID: id}

	sql, err := db.QueryParse("example/delete", arg)
	if err != nil {
		util.WriteErrorJson(w, err)
		return
	}

	_, err = h.DB.Exec(sql)
	if err != nil {
		util.WriteErrorJson(w, err)
		return
	}
}

func (h *Handler) ExampleInsert(w http.ResponseWriter, r *http.Request) {

	type InsertParams struct {
		Name string `json:"name"`
	}

	var p InsertParams

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		util.WriteErrorJson(w, err)
		return
	}

	sql, err := db.QueryParse("example/insert", p)
	if err != nil {
		util.WriteErrorJson(w, err)
		return
	}

	_, err = h.DB.Exec(sql)
	if err != nil {
		util.WriteErrorJson(w, err)
		return
	}
}
