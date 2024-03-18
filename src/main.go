package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr     string `json:"addr"`
	EchoName string `json:"echo_name"`
}

func getHandler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request", r.Method, r.URL.String())
		defer fmt.Println("response", r.Method, r.URL.String(), "Hello", name)
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func parseConfig(pathname string) (*Config, error) {
	in, err := os.Open(pathname)
	if err != nil {
		return nil, fmt.Errorf("open config file %q error: %w", pathname, err)
	}
	defer in.Close()

	var ret Config
	if err := json.NewDecoder(in).Decode(&ret); err != nil {
		return nil, fmt.Errorf("parse config file %q error: %w", pathname, err)
	}

	return &ret, nil
}

func main() {
	cfgfile := flag.String("c", "", "Config file")
	flag.Parse()

	if *cfgfile == "" {
		fmt.Println("Usage: -c <path to config.json>")
		return
	}

	cfg, err := parseConfig(*cfgfile)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", getHandler(cfg.EchoName))
	if err := http.ListenAndServe(cfg.Addr, nil); err != nil {
		log.Fatal("listen and serve error:", err)
	}
}
