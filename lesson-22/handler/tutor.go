package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"lesson22/model"
	"net/http"
)

func (h *Handler) CreateTutor(w http.ResponseWriter, r *http.Request) {
	var tutor model.Tutor
	b, err := io.ReadAll(r.Body)
	defer func() { r.Body.Close() }()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &tutor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.tutorRepo.CreateTutor(tutor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "successfully created")
}

func (h *Handler) GetTutor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	tutor, err := h.tutorRepo.GetTutor(id)

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
	err = json.NewEncoder(w).Encode(tutor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) GetAllTutors(w http.ResponseWriter, r *http.Request) {
	tutors, err := h.tutorRepo.GetAllTutors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tutorResponse := model.TutorResponse{Data: tutors}
	tutorsByte, err := json.Marshal(tutorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(tutorsByte)
}

func (h *Handler) UpdateTutor(w http.ResponseWriter, r *http.Request) {
	var tutor model.Tutor
	b, err := io.ReadAll(r.Body)
	defer func() { r.Body.Close() }()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &tutor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.tutorRepo.UpdateTutor(tutor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "successfully updated")
}

func (h *Handler) DeleteTutor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.tutorRepo.DeleteTutor(id)

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
