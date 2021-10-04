package repository

import "alexis.zapata-github.com/capstone-project/domain/model"

type PokemonRepository interface {
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
	FindById(p *model.Pokemon, id string) (*model.Pokemon, error)
}
