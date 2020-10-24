package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPokedex(t *testing.T) {
	t.Run("returns 1st pokemon name", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/pokemon/1", nil)
		response := httptest.NewRecorder()

		PokedexServer(response, request)

		got := response.Body.String()
		want := "Bulbassaur"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("returns 2nd pokemon name", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/pokemon/2", nil)
		response := httptest.NewRecorder()

		PokedexServer(response, request)

		got := response.Body.String()
		want := "Ivysaur"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
