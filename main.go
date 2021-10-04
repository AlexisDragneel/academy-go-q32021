package main

import (
	"alexis.zapata-github.com/capstone-project/infrastructure/router"
	"alexis.zapata-github.com/capstone-project/registry"
	"fmt"
	"github.com/labstack/echo/v4"
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
