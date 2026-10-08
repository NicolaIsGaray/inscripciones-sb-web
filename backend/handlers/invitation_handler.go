package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"servidor-angular/models"
	"servidor-angular/repository"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// InvitationHandler handles HTTP requests for invitations
type InvitationHandler struct {
	studentRepo *repository.StudentRepository
}

// NewInvitationHandler creates a new invitation handler
func NewInvitationHandler() *InvitationHandler {
	return &InvitationHandler{
		studentRepo: repository.NewStudentRepository(),
	}
}

// GetAllInvitations handles GET /api/invitations
func (h *InvitationHandler) GetAllInvitations(w http.ResponseWriter, r *http.Request) {
	// Obtener todos los estudiantes que tienen invitaciones
	students, err := h.studentRepo.FindAll()
	if err != nil {
		http.Error(w, "Error al obtener invitaciones", http.StatusInternalServerError)
		return
	}

	// Extraer todas las invitaciones de todos los estudiantes
	var allInvitations []models.Invitation
	for _, student := range students {
		allInvitations = append(allInvitations, student.Invitations...)
	}

	if allInvitations == nil {
		allInvitations = []models.Invitation{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allInvitations)
}

// CreateInvitation handles POST /api/invitations
func (h *InvitationHandler) CreateInvitation(w http.ResponseWriter, r *http.Request) {
	var invitation models.Invitation
	if err := json.NewDecoder(r.Body).Decode(&invitation); err != nil {
		http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
		return
	}

	// Asignar estado por defecto si no se especifica
	if invitation.Status == "" {
		invitation.Status = "PENDING"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invitation)
}

// UpdateInvitation handles PUT /api/invitations/{groupId}
func (h *InvitationHandler) UpdateInvitation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId := vars["groupId"]

	var updateData struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
		return
	}

	// Actualizar el estado de la invitación en todos los estudiantes que la tengan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := h.studentRepo.GetCollection()
	filter := bson.M{"invitations.groupId": groupId}
	update := bson.M{
		"$set": bson.M{
			"invitations.$.status": updateData.Status,
		},
	}

	result, err := collection.UpdateMany(ctx, filter, update)
	if err != nil {
		http.Error(w, "Error al actualizar invitación", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "Invitación no encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Invitación actualizada",
		"count":   result.ModifiedCount,
	})
}
