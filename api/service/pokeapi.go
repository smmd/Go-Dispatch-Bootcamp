package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/smmd/academy-go-q42021/model"
)

type setter interface {
	WritePokeMonsters(response model.Response, filePath string) error
}

type WriteService struct {
	repo setter
}

func NewWriteService(repo setter) WriteService {
	return WriteService{repo}
}

// ConsumeNationalPokedex fill the db using pokeapi external api
func (ws WriteService) ConsumeNationalPokedex() error {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/1/")

	if err != nil {
		return err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var responseObject model.Response
	json.Unmarshal(responseData, &responseObject)

	err = ws.repo.WritePokeMonsters(responseObject, FileName)

	if err != nil {
		return err
	}

	return nil
}
