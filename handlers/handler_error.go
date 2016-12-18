package handlers

type HandlerError struct {
	Error string `json:"error"`
}

func NewHandlerError(e error) HandlerError {
	return HandlerError{Error: e.Error()}
}
