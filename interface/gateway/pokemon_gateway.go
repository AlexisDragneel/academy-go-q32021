package gateway

import (
	"alexis.zapata-github.com/capstone-project/domain/gateway"
	"alexis.zapata-github.com/capstone-project/domain/model"
	"alexis.zapata-github.com/capstone-project/interface/context"
	"encoding/json"
	"fmt"
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

func NewPokemonGateway() PokemonGateway {
	return &pokemonGateway{}
}

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
