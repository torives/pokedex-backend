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

	name := PokemonName(index)

	fmt.Fprint(w, name)
}

func PokemonName(index string) string {
	if index == "1" {
		return "Bulbassaur"
	} else if index == "2" {
		return "Ivysaur"
	}

	return ""
}
