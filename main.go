package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	handler := http.HandlerFunc(PokedexServer)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Could not listen on port 8080. %v", err)
	}
}

func PokedexServer(w http.ResponseWriter, r *http.Request) {
	index := strings.TrimPrefix(r.URL.Path, "/pokemon/")

	if index == "1" {
		fmt.Fprint(w, "Bulbassaur")
		return
	}

	if index == "2" {
		fmt.Fprint(w, "Ivysaur")
		return
	}
}
