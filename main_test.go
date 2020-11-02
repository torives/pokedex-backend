package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPokedex(t *testing.T) {
	store := StubPokemonStore{
		map[string]string{
			"1": "Bulbassaur",
			"2": "Ivysaur",
		},
	}
	server := NewPokedexServer(&store)

	t.Run("returns 1st pokemon name", func(t *testing.T) {
		request := newGetPokemonRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "Bulbassaur")
	})

	t.Run("returns 2nd pokemon name", func(t *testing.T) {
		request := newGetPokemonRequest("2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "Ivysaur")
	})

	t.Run("returns 404 on invalid pokemon", func(t *testing.T) {
		request := newGetPokemonRequest("9999")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusNotFound)
	})

	t.Run("returns 200 on /pokemons", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/pokemons", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
	})
}

func TestIntegration(t *testing.T) {
	store := NewInMemoryPokemonStore()
	server := NewPokedexServer(store)
	index := "1"
	response := httptest.NewRecorder()

	server.ServeHTTP(response, newGetPokemonRequest(index))

	assertStatusCode(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "Bulbassaur")
}

func newGetPokemonRequest(index string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/pokemons/%s", index), nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func assertStatusCode(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("did not get correct status code. Got %d, wanted %d", got, want)
	}
}

type StubPokemonStore struct {
	scores map[string]string
}

func (s *StubPokemonStore) PokemonName(index string) string {
	name := s.scores[index]
	return name
}
