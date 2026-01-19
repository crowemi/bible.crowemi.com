package handlers

import (
	"net/http"
)

type Handler interface {
	GetMany(w http.ResponseWriter, r *http.Request)
}
