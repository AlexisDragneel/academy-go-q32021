package router

import (
	"github.com/AlexisDragneel/academy-go-q3202/interface/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pokemons", func(context echo.Context) error { return c.Pokemon.GetPokemons(context) })
	e.GET("/pokemons/async", func(context echo.Context) error { return c.Pokemon.GetAsyncPokemons(context) })
	e.GET("/pokemons/:id", func(context echo.Context) error { return c.Pokemon.GetPokemonById(context) })
	e.POST("/pokemons", func(context echo.Context) error { return c.Pokemon.PostPokemons(context) })
	return e
}
