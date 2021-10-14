package gateway

import "github.com/AlexisDragneel/academy-go-q3202/domain/model"

type PokemonGatewayResponse struct {
	Count   int64           `json:"count"`
	Results []model.Pokemon `json:"results"`
}
