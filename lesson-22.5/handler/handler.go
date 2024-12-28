package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func (h *Handler) InitRoutes() error {
	r := mux.NewRouter()

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("POST")
	return http.ListenAndServe(":8080", r)
}
