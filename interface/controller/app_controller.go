package controller

// AppController declaration of the main controller fo the app
type AppController struct {
	Pokemon interface{ PokemonController }
}
