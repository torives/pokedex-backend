package main

import (
	"fmt"
	"net/http"
	"strings"
)

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
