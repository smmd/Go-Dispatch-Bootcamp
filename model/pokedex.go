package model

type Response struct {
	Name            string            `json:"name"`
	PokemonOriginal []PokemonOriginal `json:"pokemon_entries"`
}

type PokemonOriginal struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}
