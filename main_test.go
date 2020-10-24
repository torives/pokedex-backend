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

		assertResponseBody(t, response.Body.String(), "Bulbassaur")
	})

	t.Run("returns 2nd pokemon name", func(t *testing.T) {
		request := newGetPokemonRequest("2")
		response := httptest.NewRecorder()

		PokedexServer(response, request)

		assertResponseBody(t, response.Body.String(), "Ivysaur")
	})
}

func newGetPokemonRequest(index string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/pokemon/%s", index), nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
