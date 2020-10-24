package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	server := new(PokedexServer)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Could not listen on port 8080. %v", err)
	}
}

type PokemonStore interface {
	PokemonName(index string) string
}

type PokedexServer struct {
	store PokemonStore
}

func (s *PokedexServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	index := strings.TrimPrefix(r.URL.Path, "/pokemon/")
	fmt.Fprint(w, s.store.PokemonName(index))
}
