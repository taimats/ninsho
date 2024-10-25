package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

type config struct {
	env string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
}

func main() {
	var cfg config
	//コマンドラインから環境変数を読み込む
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.Parse()

	fmt.Println("Running")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := AppStatus{
			Status:      "Available",
			Environment: cfg.env,
		}

		js, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	})

	err := http.ListenAndServe(":4000", corsMiddleware(http.DefaultServeMux))
	if err != nil {
		log.Println(err)
	}

}
