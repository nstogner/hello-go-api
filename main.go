package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

const message = "Hello "

	var cfg struct {
		Addr string `envconfig:"ADDR" default:":8080"`

		HTTPS bool `envconfig:"HTTPS"`

		Cert string `envconfig:"CERT" default:"server.crt"`
		Key  string `envconfig:"KEY" default:"server.key"`
		Name string `envconfig:"NAME" default:"there"`
	}

func main() {
	envconfig.MustProcess("", &cfg)

	log.Printf("starting to listen on addr %v", cfg.Addr)

	if cfg.HTTPS {
		log.Fatal(http.ListenAndServeTLS(cfg.Addr, cfg.Cert, cfg.Key, http.HandlerFunc(handle)))
	} else {
		log.Fatal(http.ListenAndServe(cfg.Addr, http.HandlerFunc(handle)))
	}
}

type response struct {
	Msg string `json:"msg"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Msg: message + cfg.Name + "!"})
}
