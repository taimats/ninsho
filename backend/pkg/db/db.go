package db

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"
)

func QueryParse(path string, data any) (q string, err error) {
	filename := "/app/sql/" + path + ".sql"

	queryTemplate, err := template.New("customFuncs").Funcs(template.FuncMap{
		"timeParse": timeParse,
		"str":       str,
	}).ParseFiles(filename)

	if err != nil {
		return "", errors.New("failed to parse query file")
	}

	queryTemplate.Templates()
	f := filepath.Base(filename)

	var result bytes.Buffer

	err = queryTemplate.ExecuteTemplate(&result, f, data)
	if err != nil {
		return "", errors.New("failed to execute query template")
	}

	q = result.String()
	queryLog(q, path)
	return q, nil

}

func queryLog(query, path string) {
	result := strings.Replace(query, "\n", "", -1)
	fmt.Println("【ログ: exec Query ▶】", path+":"+result)
}
