package handlers

import (
	"encoding/json"
	"net/http"

	"servidor-angular/repository"
)

// SettingHandler handles HTTP requests for the global settings
type SettingHandler struct {
	repo *repository.SettingRepository
}

// NewSettingHandler creates a new setting handler
func NewSettingHandler() *SettingHandler {
	return &SettingHandler{
		repo: repository.NewSettingRepository(),
	}
}

// GetSettings handles GET /api/settings (público).
// Solo expone si la formación de grupos está habilitada.
func (h *SettingHandler) GetSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := h.repo.Get()
	if err != nil {
		http.Error(w, "Error al obtener la configuración", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(settings)
}

// updateSettingsRequest representa la actualización parcial de la configuración.
// Se usa un puntero para distinguir "no enviado" de "false".
type updateSettingsRequest struct {
	GroupsEnabled *bool `json:"groupsEnabled"`
}

// UpdateSettings handles PATCH /api/admin/settings (requiere JWT)
func (h *SettingHandler) UpdateSettings(w http.ResponseWriter, r *http.Request) {
	var req updateSettingsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
		return
	}

	if req.GroupsEnabled == nil {
		http.Error(w, "Falta el campo groupsEnabled", http.StatusBadRequest)
		return
	}

	settings, err := h.repo.SetGroupsEnabled(*req.GroupsEnabled)
	if err != nil {
		http.Error(w, "Error al actualizar la configuración", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(settings)
}