package repository

import (
	"os"
	"strconv"

	"encoding/csv"

	"github.com/smmd/academy-go-q42021/model"
)

type AllPokeMonsters struct{}

type PokeMonstersWriter struct{}

func NewAllPokeMonsters() AllPokeMonsters {
	return AllPokeMonsters{}
}

func NewPokeMonstersWriter() PokeMonstersWriter {
	return PokeMonstersWriter{}
}

func (AllPokeMonsters) GetAllPokeMonsters(filePath string) (model.PokeMonsters, error) {
	lines, err := csvToObject(filePath)

	pokemons := []model.Pokemon{}
	pokeMonsters := model.PokeMonsters{pokemons}

	if err != nil {
		return pokeMonsters, err
	}

	for _, line := range lines {
		pokemon := model.Pokemon{line[0], line[1]}

		pokeMonsters.AddPokemon(pokemon)
	}

	return pokeMonsters, nil
}

func (PokeMonstersWriter) WritePokeMonsters(response model.Response, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, pokemon := range response.PokemonOriginal {
		poke := []string{strconv.Itoa(pokemon.EntryNo), pokemon.Species.Name}

		err := writer.Write(poke)

		if err != nil {
			return err
		}
	}

	return nil
}

func csvToObject(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return lines, nil
}
