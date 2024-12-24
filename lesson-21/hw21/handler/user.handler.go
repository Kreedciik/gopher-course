package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"leeson21/hw21/model"
	"net/http"
	"strconv"
)

type UserHandler struct {
	users []model.User
}

func CreateUserHandler() *UserHandler {
	return &UserHandler{users: []model.User{}}
}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	user := model.User{1, "Ilfat", 26}
	w.Header().Add("Content-Type", "application/json")
	if userId != "" {
		u, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Could not parse to json")
		}
		w.Write(u)
	}

}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	pagination := r.URL.Query().Get("pagination")
	size := r.URL.Query().Get("size")
	search := r.URL.Query().Get("search")

	filter := model.Filter{pagination, size, search, h.users}

	w.Header().Add("Content-Type", "application/json")
	f, err := json.Marshal(&filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not parse to json")
	}
	w.Write(f)

}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	body, err := io.ReadAll(r.Body)
	defer func() {
		r.Body.Close()
	}()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not get user")
		return
	}
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not parse json: %s", err.Error())
		return
	}

	h.users = append(h.users, newUser)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "user created")

}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser model.User
	body, err := io.ReadAll(r.Body)
	defer func() {
		r.Body.Close()
	}()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not get user")
		return
	}
	err = json.Unmarshal(body, &updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not parse json: %s", err.Error())
		return
	}

	for i, user := range h.users {
		if user.Id == updatedUser.Id {
			h.users[i] = updatedUser
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "user updated successfully")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "user not found")
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
	}

	for i, user := range h.users {
		if user.Id == userId {
			h.users = append(h.users[:i], h.users[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "successfully deleted")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "user not found")

}
