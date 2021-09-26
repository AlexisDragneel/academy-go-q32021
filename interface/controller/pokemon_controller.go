package controller

import (
	"alexis.zapata-github.com/capstone-project/domain/model"
	"alexis.zapata-github.com/capstone-project/interface/context"
	"alexis.zapata-github.com/capstone-project/interface/gateway"
	"alexis.zapata-github.com/capstone-project/usecase/interactor"
	"alexis.zapata-github.com/capstone-project/utils/errors"
	"net/http"
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
	var p *model.Pokemon
	id := c.Param("id")

	p, err := pc.pokemonInteractor.GetById(p, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if p == nil {
		response := &errors.HttpError{}
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
