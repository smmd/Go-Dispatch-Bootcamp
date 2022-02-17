package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smmd/academy-go-q42021/model"
)

var expectedPokedex = model.PokeMonsters{
	[]model.Pokemon{
		{
			ID:   "1",
			Name: "bulbasaur",
		},
		{
			ID:   "2",
			Name: "ivysaur",
		},
		{
			ID:   "3",
			Name: "venusaur",
		},
	},
}

func TestAllPokeMonsters_GetAllPokeMonsters(t *testing.T) {
	testCases := []struct {
		name     string
		expected model.PokeMonsters
		hasError bool
		err      error
		argument string
	}{
		{
			"getall pokemons properly",
			expectedPokedex,
			false,
			nil,
			"fixtures/pokedex_test.csv",
		},
		{
			"file no exist error",
			model.PokeMonsters{[]model.Pokemon{}},
			true,
			errors.New("open fixtures/pokedex_data_fail.csv: no such file or directory"),
			"fixtures/pokedex_data_fail.csv",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := NewAllPokeMonsters()

			actual, err := repo.GetAllPokeMonsters(tc.argument)

			assert.Equal(t, actual, tc.expected)
			if tc.hasError {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCsvRepo_WritePokeMonsters(t *testing.T) {
	response := model.Response{
		"jotto",
		[]model.PokemonOriginal{
			{
				1,
				model.PokemonSpecies{
					"bulbasaur",
				},
			},
		},
	}

	repo := NewPokeMonstersWriter()

	err := repo.WritePokeMonsters(response, "fixtures/pokedex_test.csv")

	if err != nil {
		assert.EqualError(t, err, err.Error())
	} else {
		assert.Nil(t, err)
	}
}
