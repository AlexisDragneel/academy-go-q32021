package router

import (
	"alexis.zapata-github.com/capstone-project/interface/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pokemons", func(context echo.Context) error { return c.Pokemon.GetPokemons(context) })
	e.GET("/pokemons/:id", func(context echo.Context) error { return c.Pokemon.GetPokemonById(context) })
	return e
}
