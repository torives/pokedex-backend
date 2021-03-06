package main

type PokemonStore interface {
	PokemonName(index string) string
	GetPokemonList() []Pokemon
}

type InMemoryPokemonStore struct {
	store map[string]string
}

func NewInMemoryPokemonStore() *InMemoryPokemonStore {
	return &InMemoryPokemonStore{
		map[string]string{
			"1": "Bulbasaur",
			"2": "Ivysaur",
		},
	}
}

func (i *InMemoryPokemonStore) PokemonName(index string) string {
	return i.store[index]
}

func (i *InMemoryPokemonStore) GetPokemonList() []Pokemon {

}
