package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPokedex(t *testing.T) {
	store := StubPokemonStore{
		map[string]string{
			"1": "Bulbasaur",
			"2": "Ivysaur",
		},
		nil,
	}
	server := NewPokedexServer(&store)

	t.Run("returns 1st pokemon name", func(t *testing.T) {
		request := newGetPokemonRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "Bulbasaur")
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

	t.Run("returns list of pokemons as JSON", func(t *testing.T) {

		want := []Pokemon{
			{"Bulbasaur"},
			{"Ivysaur"},
		}

		store = StubPokemonStore{nil, want}

		request, _ := http.NewRequest(http.MethodGet, "/pokemons", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Pokemon

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Pokemon, '%v'", response.Body, err)
		}

		assertStatusCode(t, response.Code, http.StatusOK)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestIntegration(t *testing.T) {
	store := NewInMemoryPokemonStore()
	server := NewPokedexServer(store)
	index := "1"
	response := httptest.NewRecorder()

	server.ServeHTTP(response, newGetPokemonRequest(index))

	assertStatusCode(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "Bulbasaur")
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
	pokemons    map[string]string
	pokemonList []Pokemon
}

func (s *StubPokemonStore) PokemonName(index string) string {
	name := s.pokemons[index]
	return name
}
