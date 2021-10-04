package controller

import (
	"github.com/AlexisDragneel/academy-go-q3202/domain/model"
	"github.com/AlexisDragneel/academy-go-q3202/interface/context"
	"github.com/AlexisDragneel/academy-go-q3202/interface/gateway"
	"github.com/AlexisDragneel/academy-go-q3202/usecase/interactor"
	"github.com/AlexisDragneel/academy-go-q3202/utils"
	"net/http"
	"strconv"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
	pokemonGateway    gateway.PokemonGateway
}

type PokemonController interface {
	GetPokemons(c context.Context) error
	GetPokemonById(c context.Context) error
	PostPokemons(c context.Context) error
}

func NewPokemonController(pi interactor.PokemonInteractor, pg gateway.PokemonGateway) PokemonController {
	return &pokemonController{pi, pg}
}

func (pc *pokemonController) GetPokemons(c context.Context) error {
	var p []*model.Pokemon

	p, err := pc.pokemonInteractor.Get(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, p)
}

func (pc *pokemonController) GetPokemonById(c context.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		return err
	}

	p := &model.Pokemon{
		ID: id,
	}

	p, err = pc.pokemonInteractor.GetById(p)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if p == nil {
		response := &utils.HttpError{}
		return c.JSON(http.StatusNotFound, response.Fill404Error())
	}

	return c.JSON(http.StatusOK, p)
}

func (pc *pokemonController) PostPokemons(c context.Context) error {
	var p []*model.Pokemon

	p, err := pc.pokemonGateway.FetchPokemons(p, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	quantity, err := pc.pokemonInteractor.PostPokemons(p)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, quantity)
}
