package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PokedexServer)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Could not listen on port 8080. %v", err)
	}
}

func PokedexServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bulbassaur")
}
