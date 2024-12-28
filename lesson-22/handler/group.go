package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"lesson22/model"
	"net/http"
)

func (h *Handler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group model.StudentGroup
	b, err := io.ReadAll(r.Body)
	defer func() { r.Body.Close() }()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.groupRepo.CreateGroup(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "successfully created")
}

func (h *Handler) GetGroup(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	group, err := h.groupRepo.GetGroup(id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := h.groupRepo.GetAllGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tutorResponse := model.GroupResponse{Data: groups}
	groupsByte, err := json.Marshal(tutorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(groupsByte)
}

func (h *Handler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	var group model.StudentGroup
	b, err := io.ReadAll(r.Body)
	defer func() { r.Body.Close() }()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.groupRepo.UpdateGroup(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "successfully updated")
}

func (h *Handler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.groupRepo.DeleteGroup(id)

	if err != nil {
		if err.Error() == "not listed" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	fmt.Fprintf(w, "successfully deleted")
}
