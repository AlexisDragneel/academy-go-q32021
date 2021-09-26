package registry

import (
	"alexis.zapata-github.com/capstone-project/interface/controller"
	"alexis.zapata-github.com/capstone-project/interface/gateway"
	ir "alexis.zapata-github.com/capstone-project/interface/repository"
	"alexis.zapata-github.com/capstone-project/usecase/interactor"
	ur "alexis.zapata-github.com/capstone-project/usecase/repository"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor(), r.NewPokemonGateway())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository())
}

func (r *registry) NewPokemonRepository() ur.PokemonRepository {
	return ir.NewPokemonRepository()
}

func (r *registry) NewPokemonGateway() gateway.PokemonGateway {
	return gateway.NewPokemonGateway()
}
