package router

import (
	"github.com/gin-gonic/gin"

	"github.com/smmd/go-dispatch-bootcamp/api/service"
	"github.com/smmd/go-dispatch-bootcamp/controller"
	"github.com/smmd/go-dispatch-bootcamp/repository"
	"github.com/smmd/go-dispatch-bootcamp/wpool"
)

func Route() {
	searchService := service.NewSearchService(repository.NewAllPokeMonsters())
	apiService := service.NewWriteService(repository.NewPokeMonstersWriter())
	worker := wpool.NewPokemonWorker()

	apiController := controller.NewPokemonsHandler(searchService, apiService, worker)

	router := gin.Default()
	router.GET("/pokemonsters/", apiController.PokeMonsters)
	router.GET("/pokemonsters/:id", apiController.Pokemon)
	router.GET("/fill-pokedex/", apiController.Pokedex)

	router.GET("/worker/:type/:items/:items_per_workers",
		apiController.PokeMonstersByWorker)

	router.Run(":3001")
}
