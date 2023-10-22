// This is the handler package, where you can implement your API endpoints.
// This file contains the handler struct, which implements the ServerInterface
// You can inject dependencies here, such as repositories or interactors / use case.

package handler

type Handler struct {
}

// New returns a new Handler.
// TODO: You may inject dependencies here, such as repositories or interactors / use cases.
func New() *Handler {
	return &Handler{}
}
