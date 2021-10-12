package controller

import (
	"github.com/AlexisDragneel/academy-go-q3202/domain/model"
	"github.com/AlexisDragneel/academy-go-q3202/interface/context"
	"github.com/AlexisDragneel/academy-go-q3202/interface/gateway"
	"github.com/AlexisDragneel/academy-go-q3202/usecase/interactor"
	"github.com/AlexisDragneel/academy-go-q3202/utils"
	"net/http"
	"strconv"
	"strings"
)

const (
	itemsParam       = "items"
	itemsWorkerParam = "items_per_workers"
	typeParam        = "type"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
	pokemonGateway    gateway.PokemonGateway
}

type PokemonController interface {
	GetPokemons(c context.Context) error
	GetAsyncPokemons(c context.Context) error
	GetPokemonById(c context.Context) error
	PostPokemons(c context.Context) error
}

// NewPokemonController creates a new instance of the pokemon controller to manage the dependency injections
func NewPokemonController(pi interactor.PokemonInteractor, pg gateway.PokemonGateway) PokemonController {
	return &pokemonController{pi, pg}
}

// GetPokemons function that communicates to the interactor for getting all the pokemons for the csv file
func (pc *pokemonController) GetPokemons(c context.Context) error {
	var p []*model.Pokemon

	p, err := pc.pokemonInteractor.Get(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}

// GetAsyncPokemons function that returns the list of pokemons with worker pool
func (pc *pokemonController) GetAsyncPokemons(c context.Context) error {
	var p []*model.Pokemon

	items, err := strconv.ParseInt(c.QueryParam(itemsParam), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, "query param 'items' must contain a value and  should be number"))
	}

	itemsWorker, err := strconv.ParseInt(c.QueryParam(itemsWorkerParam), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, "query param 'items_per_workers' must contain a value and  should be number"))
	}

	if itemsWorker > items {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, "items_per_workers couldn't be bigger than items"))
	}

	t := strings.ToLower(c.QueryParam(typeParam))
	if t != "" && strings.Compare(t, utils.Odd) != 0 && strings.Compare(t, utils.Even) != 0 {
		return c.JSON(http.StatusBadRequest, utils.CreateResponse(http.StatusBadRequest, "query param 'type' only supports 'even' and 'odd'"))
	}

	p, err = pc.pokemonInteractor.GetAsync(p, t, items, itemsWorker)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.CreateResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, p)
}

// GetPokemonById function that communicates to the interactor for getting a pokemon based on the id received
func (pc *pokemonController) GetPokemonById(c context.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.CreateResponse(http.StatusInternalServerError, err.Error()))
	}

	p := &model.Pokemon{
		ID: id,
	}

	p, err = pc.pokemonInteractor.GetById(p)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.CreateResponse(http.StatusInternalServerError, err.Error()))
	}

	if p == nil {
		return c.JSON(http.StatusNotFound, utils.CreateResponse(http.StatusNotFound, "Item Not Found"))
	}

	return c.JSON(http.StatusOK, p)
}

// PostPokemons function that calls a gateway to get information and call the interactor for append the information to the csv file
func (pc *pokemonController) PostPokemons(c context.Context) error {
	var p []*model.Pokemon

	p, err := pc.pokemonGateway.FetchPokemons(p, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.CreateResponse(http.StatusInternalServerError, err.Error()))
	}

	quantity, err := pc.pokemonInteractor.PostPokemons(p)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.CreateResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, quantity)
}
