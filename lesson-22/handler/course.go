package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"lesson22/model"
	"net/http"
)

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course model.Course
	b, err := io.ReadAll(r.Body)
	defer func() { r.Body.Close() }()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.courseRepo.CreateCourse(course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "successfully created")
}

func (h *Handler) GetCourse(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	course, err := h.courseRepo.GetCourse(id)

	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	var course model.Course
	b, err := io.ReadAll(r.Body)
	defer func() { r.Body.Close() }()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.courseRepo.UpdateCourse(course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "successfully updated")
}

func (h *Handler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.courseRepo.DeleteCourse(id)

	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	fmt.Fprintf(w, "successfully deleted")
}
