package handlers

import (
	"encoding/json"
	"net/http"

	"servidor-angular/models"
	"servidor-angular/repository"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// StudentHandler handles HTTP requests for students
type StudentHandler struct {
	repo *repository.StudentRepository
}

// NewStudentHandler creates a new student handler
func NewStudentHandler() *StudentHandler {
	return &StudentHandler{
		repo: repository.NewStudentRepository(),
	}
}

// GetAllStudents handles GET /api/students
func (h *StudentHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.repo.FindAll()
	if err != nil {
		http.Error(w, "Error al obtener estudiantes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// GetStudent handles GET /api/students/{id}
func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	student, err := h.repo.FindByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al obtener estudiante", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// CreateStudent handles POST /api/students
func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&student); err != nil {
		http.Error(w, "Error al crear estudiante", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

// UpdateStudent handles PUT /api/students/{id}
func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(id, &student); err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al actualizar estudiante", http.StatusInternalServerError)
		return
	}

	student.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// DeleteStudent handles DELETE /api/students/{id}
func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al eliminar estudiante", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
