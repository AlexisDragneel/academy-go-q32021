package registry

import (
	"github.com/AlexisDragneel/academy-go-q3202/interface/controller"
	"github.com/AlexisDragneel/academy-go-q3202/interface/gateway"
	ir "github.com/AlexisDragneel/academy-go-q3202/interface/repository"
	"github.com/AlexisDragneel/academy-go-q3202/usecase/interactor"
	ur "github.com/AlexisDragneel/academy-go-q3202/usecase/repository"
)

// NewPokemonController initializer for the pokemon controller inside the registry
func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor(), r.NewPokemonGateway())
}

// NewPokemonInteractor initializer of the pokemon interactor inside the registry
func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository())
}

// NewPokemonRepository initializer of the pokemon repository in the registry
func (r *registry) NewPokemonRepository() ur.PokemonRepository {
	return ir.NewPokemonRepository()
}

// NewPokemonGateway initializer of the pokemon gateway inside the registry
func (r *registry) NewPokemonGateway() gateway.PokemonGateway {
	return gateway.NewPokemonGateway()
}
