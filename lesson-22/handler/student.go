package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"lesson22/model"
	"net/http"
)

func (h *Handler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student model.CreateStudentRequest
	b, err := io.ReadAll(r.Body)
	defer func() { r.Body.Close() }()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.studentRepo.CreateStudent(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "successfully created")
}
func (h *Handler) GetStudent(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	student, err := h.studentRepo.GetStudent(id)

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
	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (h *Handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var student model.Student
	b, err := io.ReadAll(r.Body)
	defer func() { r.Body.Close() }()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.studentRepo.UpdateStudent(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "successfully updated")
}
func (h *Handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.studentRepo.DeleteStudent(id)

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

func (h *Handler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.studentRepo.GetAllStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	studentResponse := model.StudentResponse{Data: students}
	studentsByte, err := json.Marshal(studentResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(studentsByte)
}
