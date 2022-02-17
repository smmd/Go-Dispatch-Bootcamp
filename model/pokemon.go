package model

type Pokemon struct {
	ID   string `json:"pokemon_id"`
	Name string `json:"name"`
}

type PokeMonsters struct {
	Pokemons []Pokemon
}

func NewPokemon(id string, name string) *Pokemon {
	p := new(Pokemon)
	p.ID = id
	p.Name = name

	return p
}

func (pokedex *PokeMonsters) AddPokemon(pokemon Pokemon) []Pokemon {
	pokedex.Pokemons = append(pokedex.Pokemons, pokemon)

	return pokedex.Pokemons
}

func (p Pokemon) GetId() string {
	return p.ID
}

func (p Pokemon) GetName() string {
	return p.Name
}
