package main

type PokemonStore interface {
	PokemonName(index string) string
}

type InMemoryPokemonStore struct {
	store map[string]string
}

func NewInMemoryPokemonStore() *InMemoryPokemonStore {
	return &InMemoryPokemonStore{
		map[string]string{
			"1": "Bulbassaur",
			"2": "Ivysaur",
		},
	}
}

func (i *InMemoryPokemonStore) PokemonName(index string) string {
	return i.store[index]
}
