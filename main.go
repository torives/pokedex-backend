package main

import (
	"fmt"
	"net/http"
)

func main() {
	println("hello!")
}

func PokedexServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bulbassaur")
}
