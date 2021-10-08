package utils

import "net/http"

type HttpError struct {
	Code    int
	Message string
}

func (h *HttpError) Fill404Error() *HttpError {

	if h.Code == 0 {
		h.Code = http.StatusNotFound
	}

	if h.Message == "" {
		h.Message = "Item Not Found"
	}

	return h
}
