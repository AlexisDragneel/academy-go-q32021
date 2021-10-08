package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/AlexisDragneel/academy-go-q3202/infrastructure/router"
	"github.com/AlexisDragneel/academy-go-q3202/registry"
)

func main() {

	r := registry.NewRegistry()

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost:8000")
	if err := e.Start(":8000"); err != nil {
		fmt.Printf("unable to start server: %v", err)
	}

}
