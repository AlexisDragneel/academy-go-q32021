package context

// Context extraction of the interface echo.Context from the library of echo
type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(name string) string
	QueryParam(name string) string
}
