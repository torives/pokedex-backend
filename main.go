package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	server := &PokedexServer{&InMemoryPokemonStore{}}

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

	name := s.store.PokemonName(index)

	if name == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, name)
}

type InMemoryPokemonStore struct{}

func (i *InMemoryPokemonStore) PokemonName(index string) string {
	return "oi"
}
