package main

import (
	"backend/cmd/api/handler"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type config struct {
	env string
	db  struct {
		dsn string //dsnはデータソースネーム
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
}

func main() {
	var cfg config
	//コマンドラインから環境変数を読み込む
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://postgres:secret@postgres/hands_on?sslmode=disable", "Postgres connection string")
	flag.Parse()

	fmt.Println("Running")

	db, err := OpenDB(cfg)
	if err != nil {
		log.Fatalf("Unable to open database: %v", err)
	}
	defer db.Close()

	routes(&handler.Handler{DB: db})

	err = http.ListenAndServe(":4000", corsMiddleware(http.DefaultServeMux))
	if err != nil {
		log.Println(err)
	}

}

func OpenDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
