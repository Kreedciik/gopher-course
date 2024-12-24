package main

import (
	"fmt"
	"leeson21/hw21/handler"
	"log"
	"net/http"
)

func readAuthorization(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if len(token) > 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, token)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Token was not provided by Authorization header")
	}
}

func main() {

	h := handler.CreateUserHandler()
	http.HandleFunc("POST /readBearer", readAuthorization)
	http.HandleFunc("GET /getUser/{id}", h.GetUser)
	http.HandleFunc("GET /getUsers", h.GetUsers)
	http.HandleFunc("POST /user/create", h.CreateUser)
	http.HandleFunc("PATCH /user/update", h.UpdateUser)
	http.HandleFunc("DELETE /user/delete/{id}", h.DeleteUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
