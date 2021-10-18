package utils

type HttpResponse struct {
	Code    int
	Message string
}

func CreateResponse(code int, message string) HttpResponse {
	return HttpResponse{
		code,
		message,
	}
}
