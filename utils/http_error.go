package utils

type HttpError struct {
	Code    int
	Message string
}

func CreateError(code int, message string) HttpError {
	return HttpError{
		code,
		message,
	}
}
