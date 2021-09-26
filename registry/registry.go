package registry

import "github.com/AlexisDragneel/academy-go-q3202/interface/controller"

type registry struct {
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Pokemon: r.NewPokemonController(),
	}
}
