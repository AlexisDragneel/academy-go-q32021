package interactor

import (
	"alexis.zapata-github.com/capstone-project/domain/model"
	"alexis.zapata-github.com/capstone-project/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
}

type PokemonInteractor interface {
	Get(p []*model.Pokemon) ([]*model.Pokemon, error)
	GetById(p *model.Pokemon, id string) (*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository) PokemonInteractor {
	return &pokemonInteractor{r}
}

func (pi *pokemonInteractor) Get(p []*model.Pokemon) ([]*model.Pokemon, error) {

	p, err := pi.PokemonRepository.FindAll(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (pi *pokemonInteractor) GetById(p *model.Pokemon, id string) (*model.Pokemon, error) {

	p, err := pi.PokemonRepository.FindById(p, id)
	if err != nil {
		return nil, err
	}

	return p, nil
}
