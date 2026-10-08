package router

import (
	"net/http"

	"servidor-angular/handlers"

	"github.com/gorilla/mux"
)

// NewRouter creates and configures the API router
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Inicializar handlers
	studentHandler := handlers.NewStudentHandler()
	groupHandler := handlers.NewGroupHandler()
	invitationHandler := handlers.NewInvitationHandler()

	// Rutas de Students
	r.HandleFunc("/api/students", studentHandler.GetAllStudents).Methods("GET")
	r.HandleFunc("/api/students", studentHandler.CreateStudent).Methods("POST")
	r.HandleFunc("/api/students/{id}", studentHandler.GetStudent).Methods("GET")
	r.HandleFunc("/api/students/{id}", studentHandler.UpdateStudent).Methods("PUT")
	r.HandleFunc("/api/students/{id}", studentHandler.DeleteStudent).Methods("DELETE")

	// Rutas de Groups
	r.HandleFunc("/api/groups", groupHandler.GetAllGroups).Methods("GET")
	r.HandleFunc("/api/groups", groupHandler.CreateGroup).Methods("POST")
	r.HandleFunc("/api/groups/{id}", groupHandler.GetGroup).Methods("GET")
	r.HandleFunc("/api/groups/{id}", groupHandler.UpdateGroup).Methods("PUT")
	r.HandleFunc("/api/groups/{id}", groupHandler.DeleteGroup).Methods("DELETE")

	// Rutas de Invitations
	r.HandleFunc("/api/invitations", invitationHandler.GetAllInvitations).Methods("GET")
	r.HandleFunc("/api/invitations", invitationHandler.CreateInvitation).Methods("POST")
	r.HandleFunc("/api/invitations/{groupId}", invitationHandler.UpdateInvitation).Methods("PUT")

	return r
}

// CorsMiddleware handles CORS headers
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
