package controller

import (
	"alexis.zapata-github.com/capstone-project/domain/model"
	"alexis.zapata-github.com/capstone-project/usecase/interactor"
	"net/http"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemons(c Context) error
	GetPokemonById(c Context) error
}

func NewPokemonController(pi interactor.PokemonInteractor) PokemonController {
	return &pokemonController{pi}
}

func (pc *pokemonController) GetPokemons(c Context) error {
	var p []*model.Pokemon

	p, err := pc.pokemonInteractor.Get(p)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func (pc *pokemonController) GetPokemonById(c Context) error {
	var p *model.Pokemon
	id := c.Param("id")

	p, err := pc.pokemonInteractor.GetById(p, id)

	if err != nil {
		return err
	}

	if p == nil {
		return c.JSON(http.StatusNotFound, p)
	}

	return c.JSON(http.StatusOK, p)
}
