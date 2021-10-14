package gateway

import (
	"encoding/json"
	"fmt"
	"github.com/AlexisDragneel/academy-go-q3202/domain/gateway"
	"github.com/AlexisDragneel/academy-go-q3202/domain/model"
	"github.com/AlexisDragneel/academy-go-q3202/interface/context"
	"io/ioutil"
	"net/http"
	"strconv"
)

const url = "https://pokeapi.co/api/v2/pokemon"

type pokemonGateway struct {
}

type PokemonGateway interface {
	FetchPokemons(p []*model.Pokemon, c context.Context) ([]*model.Pokemon, error)
}

// NewPokemonGateway function that creates a new instance of the gateway for manage the dependency injection
func NewPokemonGateway() PokemonGateway {
	return &pokemonGateway{}
}

// FetchPokemons function that receives the manage an external api call for getting a list of pokemons
func (pg *pokemonGateway) FetchPokemons(p []*model.Pokemon, c context.Context) ([]*model.Pokemon, error) {

	offset, err := strconv.ParseUint(c.QueryParam("offset"), 10, 64)

	if err != nil {
		offset = 0
	}

	request, err := http.Get(fmt.Sprintf("%v?offset=%v", url, offset))

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		return nil, err
	}

	response := gateway.PokemonGatewayResponse{}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	for index, value := range response.Results {

		id := uint64(index+1) + offset

		pokemon := &model.Pokemon{
			ID:   id,
			Name: value.Name,
		}

		p = append(p, pokemon)

	}

	return p, nil
}
