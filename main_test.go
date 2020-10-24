package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPokedex(t *testing.T) {
	t.Run("returns 1st pokemon name", func(t *testing.T) {
		request := newGetPokemonRequest("1")
		response := httptest.NewRecorder()

		PokedexServer(response, request)

		got := response.Body.String()
		want := "Bulbassaur"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("returns 2nd pokemon name", func(t *testing.T) {
		request := newGetPokemonRequest("2")
		response := httptest.NewRecorder()

		PokedexServer(response, request)

		got := response.Body.String()
		want := "Ivysaur"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}

func newGetPokemonRequest(index string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/pokemon/%s", index), nil)
	return request
}

