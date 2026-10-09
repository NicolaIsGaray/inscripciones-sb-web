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

// GroupHandler handles HTTP requests for groups
type GroupHandler struct {
	repo *repository.GroupRepository
}

// NewGroupHandler creates a new group handler
func NewGroupHandler() *GroupHandler {
	return &GroupHandler{
		repo: repository.NewGroupRepository(),
	}
}

// GetAllGroups handles GET /api/groups
func (h *GroupHandler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := h.repo.FindAll()
	if err != nil {
		http.Error(w, "Error al obtener grupos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

// GetGroup handles GET /api/groups/{id}
func (h *GroupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	group, err := h.repo.FindByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Grupo no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al obtener grupo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

// CreateGroup handles POST /api/groups
func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group models.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&group); err != nil {
		http.Error(w, "Error al crear grupo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
}

// UpdateGroup handles PUT /api/groups/{id}
func (h *GroupHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var group models.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
		return
	}

	if err := h.repo.Update(id, &group); err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Grupo no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al actualizar grupo", http.StatusInternalServerError)
		return
	}

	group.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

// DeleteGroup handles DELETE /api/groups/{id}
func (h *GroupHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Grupo no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al eliminar grupo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RemoveMember handles DELETE /api/admin/groups/{id}/members/{memberId}
func (h *GroupHandler) RemoveMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	memberId := vars["memberId"]

	// Obtener el grupo
	group, err := h.repo.FindByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Grupo no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al obtener grupo", http.StatusInternalServerError)
		return
	}

	// Filtrar el miembro a eliminar
	var newMembers []models.Student
	for _, member := range group.Members {
		if member.ID.Hex() != memberId {
			newMembers = append(newMembers, member)
		}
	}
	group.Members = newMembers

	// Actualizar el grupo
	if err := h.repo.Update(id, group); err != nil {
		http.Error(w, "Error al actualizar grupo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}
