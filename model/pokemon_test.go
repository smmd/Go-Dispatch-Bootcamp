package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPokeMonsters_NewPokemon(t *testing.T) {
	poke := NewPokemon("3", "pikachu")
	expected := &Pokemon{
		"3",
		"pikachu",
	}

	actual := poke
	assert.Equal(t, actual, expected)
}

func TestPokeMonsters_AddPokemon(t *testing.T) {
	initial := []Pokemon{
		{
			"29",
			"nidoran-f",
		},
		{
			"30",
			"nidorina",
		},
	}

	expected := []Pokemon{
		{
			"29",
			"nidoran-f",
		},
		{
			"30",
			"nidorina",
		},
		{
			"21",
			"poke-test",
		},
	}

	pokemon := Pokemon{"21", "poke-test"}
	pokemonsters := PokeMonsters{initial}

	actual := pokemonsters.AddPokemon(pokemon)

	assert.Equal(t, actual, expected)
}
