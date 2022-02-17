package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/smmd/go-dispatch-bootcamp/model"
	"github.com/smmd/go-dispatch-bootcamp/wpool"
)

type search interface {
	GetAll() (model.PokeMonsters, error)
	GetOneByID(id string) (model.Pokemon, error)
}

type pokeapi interface {
	ConsumeNationalPokedex() error
}

type pokeworker interface {
	PokemonWorkerPool(wpool.Request) wpool.Response
}

type PokemonsHandler struct {
	searchService search
	apiService    pokeapi
	pokeWorker    pokeworker
}

func NewPokemonsHandler(search search, pokeapi pokeapi, pokeworker pokeworker) PokemonsHandler {
	return PokemonsHandler{
		search,
		pokeapi,
		pokeworker,
	}
}

func (ph PokemonsHandler) PokeMonsters(c *gin.Context) {
	pokeMonsters, err := ph.searchService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pokeMonsters)
}

func (ph PokemonsHandler) Pokemon(c *gin.Context) {
	id := c.Param("id")
	pokemon, err := ph.searchService.GetOneByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	c.JSON(http.StatusOK, pokemon)
}

func (ph PokemonsHandler) Pokedex(c *gin.Context) {
	err := ph.apiService.ConsumeNationalPokedex()

	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf(err.Error()))
		return
	}

	response := make(map[string]string)
	response["message"] = "OK"

	c.JSON(http.StatusOK, response)
}

func (ph PokemonsHandler) PokeMonstersByWorker(c *gin.Context) {
	numItems, _ := strconv.Atoi(c.Param("items"))
	itemsPerWorker, _ := strconv.Atoi(c.Param("items_per_workers"))

	wrequest, err := workerRequest(
		c.Param("type"),
		numItems,
		itemsPerWorker,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	response := ph.pokeWorker.PokemonWorkerPool(wrequest)

	if response.Err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func workerRequest(wtype string, items int, itemsPerWorker int) (wpool.Request, error) {
	v := validator.New()
	_ = v.RegisterValidation("enum", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "odd" || fl.Field().String() == "even"
	})

	request := wpool.Request{
		TypeOfJob:      wtype,
		NumberOfItems:  items,
		ItemsPerWorker: itemsPerWorker,
	}

	err := v.Struct(request)
	if err != nil {
		return request, fmt.Errorf("invalid request: %w", err)
	}

	return request, nil
}
