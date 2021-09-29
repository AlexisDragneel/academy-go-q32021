package registry

import (
	"github.com/AlexisDragneel/academy-go-q3202/interface/controller"
	"github.com/AlexisDragneel/academy-go-q3202/interface/gateway"
	ir "github.com/AlexisDragneel/academy-go-q3202/interface/repository"
	"github.com/AlexisDragneel/academy-go-q3202/usecase/interactor"
	ur "github.com/AlexisDragneel/academy-go-q3202/usecase/repository"
	"log"
	"os"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor(), r.NewPokemonGateway())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository())
}

func (r *registry) NewPokemonRepository() ur.PokemonRepository {

	file, err := os.OpenFile("db.csv", os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	return ir.NewPokemonRepository(file)
}

func (r *registry) NewPokemonGateway() gateway.PokemonGateway {
	return gateway.NewPokemonGateway()
}
