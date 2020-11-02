package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PokedexServer{NewInMemoryPokemonStore()}

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Could not listen on port 8080. %v", err)
	}
}
