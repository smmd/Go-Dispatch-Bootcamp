package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smmd/go-dispatch-bootcamp/model"
)

var pokemonsters = model.PokeMonsters{
	[]model.Pokemon{
		{
			"3",
			"pikachu",
		},
		{
			"28",
			"sandslash",
		},
		{
			"29",
			"nidoran-f",
		},
		{
			"30",
			"nidorina",
		},
	},
}

type mockCsvRepo struct {
	mock.Mock
}

func (ms mockCsvRepo) GetAllPokeMonsters(filePath string) (model.PokeMonsters, error) {
	arg := ms.Called(filePath)
	return arg.Get(0).(model.PokeMonsters), arg.Error(1)
}

func TestSearchService_GetAll(t *testing.T) {
	testCases := []struct {
		name     string
		response model.PokeMonsters
		argument string
		repoErr  error
		expected model.PokeMonsters
		err      error
	}{
		{
			"get all pokemons properly",
			pokemonsters,
			"repository/files/pokedex_data.csv",
			nil,
			pokemonsters,
			nil,
		},
		{
			"error when repository emtpy response",
			model.PokeMonsters{[]model.Pokemon{}},
			"repository/files/pokedex_data.csv",
			errors.New("test error"),
			model.PokeMonsters{[]model.Pokemon{}},
			errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := mockCsvRepo{}
			mockRepo.On("GetAllPokeMonsters", tc.argument).Return(tc.response, tc.repoErr)

			service := NewSearchService(mockRepo)
			actual, err := service.GetAll()

			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestSearchService_GetOneByID(t *testing.T) {
	expected := model.Pokemon{
		"3",
		"pikachu",
	}

	mockRepo := mockCsvRepo{}
	mockRepo.On("GetAllPokeMonsters", "repository/files/pokedex_data.csv").Return(pokemonsters, nil)

	service := NewSearchService(mockRepo)
	actual, err := service.GetOneByID("3")

	assert.Equal(t, expected, actual)
	assert.Equal(t, err, nil)
}
