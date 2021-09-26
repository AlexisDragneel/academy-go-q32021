package gateway

import "alexis.zapata-github.com/capstone-project/domain/model"

type PokemonGatewayResponse struct {
	Count   int64           `json:"count"`
	Results []model.Pokemon `json:"results"`
}
