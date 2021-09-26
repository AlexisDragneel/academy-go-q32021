package repository

import "github.com/AlexisDragneel/academy-go-q3202/domain/model"

type PokemonRepository interface {
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
	FindById(p *model.Pokemon, id string) (*model.Pokemon, error)
	PostPokemons(p []*model.Pokemon) (int, error)
}
