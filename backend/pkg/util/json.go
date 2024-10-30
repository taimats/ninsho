package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")
	wrapper := make(map[string]interface{})
	wrapper["data"] = data
	wrapper["err"] = nil
	if err != nil {
		wrapper["err"] = err.Error()
	}

	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		WriteErrorJson(w, err)
		return
	}
}

func WriteErrorJson(w http.ResponseWriter, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\033[31m【ログ: サーバーエラー!!! WriteErrorJSON panic▶】\033[0m")
		}
	}()

	fmt.Println("\033[31m【ログ: サーバーエラー!!! WriteErrorJSON panic▶】\033[0m", err)
	if err == nil {
		WriteJson(w, http.StatusInternalServerError, nil, errors.New("internal server error"))
	} else {
		WriteJson(w, http.StatusInternalServerError, nil, err)
	}
}
