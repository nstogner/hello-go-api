package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

const message = "Hello! 2"

func main() {
	var cfg struct {
		Addr string `envconfig:"ADDR" default:":8080"`
	}
	envconfig.MustProcess("", &cfg)

	log.Printf("starting to listen on addr %v", cfg.Addr)

	log.Fatal(http.ListenAndServe(cfg.Addr, http.HandlerFunc(handle)))
}

type response struct {
	Msg string `json:"msg"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Msg: message})
}
