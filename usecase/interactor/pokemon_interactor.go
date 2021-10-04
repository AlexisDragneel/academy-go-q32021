package interactor

import (
	"github.com/AlexisDragneel/academy-go-q3202/domain/model"
	"github.com/AlexisDragneel/academy-go-q3202/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
}

type PokemonInteractor interface {
	Get(p []*model.Pokemon) ([]*model.Pokemon, error)
	GetById(p *model.Pokemon) (*model.Pokemon, error)
	PostPokemons(p []*model.Pokemon) (int, error)
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

func (pi *pokemonInteractor) GetById(p *model.Pokemon) (*model.Pokemon, error) {

	p, err := pi.PokemonRepository.FindById(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (pi *pokemonInteractor) PostPokemons(p []*model.Pokemon) (int, error) {
	insertedPokemons, err := pi.PokemonRepository.PostPokemons(p)

	if err != nil {
		return 0, err
	}
	return insertedPokemons, nil
}
