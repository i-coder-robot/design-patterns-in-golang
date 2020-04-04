package Responsibility_Chain

import "strconv"

type Handler interface {
	Handler(handleID int) string
}

type handler struct {
	name      string
	next      Handler
	handlerID int
}

func NewHandler(name string, next Handler, handleID int) *handler {
	return &handler{name, next, handleID}
}

func (h *handler) Handler(handlerID int) string {
	if h.handlerID == handlerID {
		return h.name + " handled " + strconv.Itoa(handlerID)
	}
	return h.next.Handler(handlerID)
}
