package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokedexServer struct {
	store PokemonStore
	http.Handler
}

func NewPokedexServer(store PokemonStore) *PokedexServer {
	s := new(PokedexServer)

	s.store = store

	router := http.NewServeMux()
	router.Handle("/pokemons", http.HandlerFunc(s.pokemonListHandler))
	router.Handle("/pokemons/", http.HandlerFunc(s.pokemonNameHandler))

	s.Handler = router

	return s
}

func (s *PokedexServer) pokemonListHandler(w http.ResponseWriter, r *http.Request) {
	list := s.getPokemonList()
	json.NewEncoder(w).Encode(list)
}

func (s *PokedexServer) getPokemonList() []Pokemon {
	return []Pokemon{
		{"Bulbasaur"},
	}
}

func (s *PokedexServer) pokemonNameHandler(w http.ResponseWriter, r *http.Request) {
	index := r.URL.Path[len("/pokemons/"):]
	name := s.store.PokemonName(index)

	if name == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Pokemon not found")
	}

	fmt.Fprint(w, name)
}
