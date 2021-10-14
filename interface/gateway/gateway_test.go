package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlexisDragneel/academy-go-q3202/domain/model"

	"github.com/labstack/echo/v4"
)

func TestFetchPokemons(t *testing.T) {
	expectedLength := 20
	var res []*model.Pokemon
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/pokemons", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	pg := NewPokemonGateway()

	if res, _ = pg.FetchPokemons(res, c); len(res) < expectedLength {
		t.Errorf("the len of the response is %v and should be %v", len(res), expectedLength)
	}
}
