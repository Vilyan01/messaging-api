package handlers

import (
	"fmt"
)

type HandlerError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NotFoundError(t string) HandlerError {
	s := fmt.Sprintf("%v not found", t)
	return HandlerError{Code: 404, Message: s}
}
