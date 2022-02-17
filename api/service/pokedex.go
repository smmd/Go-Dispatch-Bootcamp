package service

import (
	"github.com/smmd/academy-go-q42021/model"
)

const FileName = "repository/files/pokedex_data.csv"

type getter interface {
	GetAllPokeMonsters(filePath string) (model.PokeMonsters, error)
}

type SearchService struct {
	repo getter
}

func NewSearchService(repo getter) SearchService {
	return SearchService{repo}
}

// GetAll returns all pokemons from db
func (ss SearchService) GetAll() (model.PokeMonsters, error) {
	return ss.repo.GetAllPokeMonsters(FileName)
}

// GetOneByID return the pokemon from db that matches the ID
func (ss SearchService) GetOneByID(id string) (model.Pokemon, error) {
	pokeMonsters, err := ss.repo.GetAllPokeMonsters(FileName)

	if err == nil {
		for _, poke := range pokeMonsters.Pokemons {
			if poke.ID == id {
				return poke, nil
			}
		}
	}

	return model.Pokemon{}, err
}
