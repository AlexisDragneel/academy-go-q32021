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
	GetAsync(p []*model.Pokemon, t string, items, itemsWorker int64) ([]*model.Pokemon, error)
	GetById(p *model.Pokemon) (*model.Pokemon, error)
	PostPokemons(p []*model.Pokemon) (int, error)
}

// NewPokemonInteractor function that returns a new instance of the interactor for manage dependency injection
func NewPokemonInteractor(r repository.PokemonRepository) PokemonInteractor {
	return &pokemonInteractor{r}
}

// Get function that communicates to the repository to fetch all the data in the DB
func (pi *pokemonInteractor) Get(p []*model.Pokemon) ([]*model.Pokemon, error) {

	p, err := pi.PokemonRepository.FindAll(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (pi *pokemonInteractor) GetAsync(p []*model.Pokemon, t string, items, itemsWorker int64) ([]*model.Pokemon, error) {

	p, err := pi.PokemonRepository.FindAllAsync(p, t, items, itemsWorker)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// GetById function that communicates with the repository to get an specific pokemon based on the id
func (pi *pokemonInteractor) GetById(p *model.Pokemon) (*model.Pokemon, error) {

	p, err := pi.PokemonRepository.FindById(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// PostPokemons function tha communicates with the repository to attach new pokemons on the csv file
func (pi *pokemonInteractor) PostPokemons(p []*model.Pokemon) (int, error) {
	insertedPokemons, err := pi.PokemonRepository.PostPokemons(p)

	if err != nil {
		return 0, err
	}
	return insertedPokemons, nil
}
