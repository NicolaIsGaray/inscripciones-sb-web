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

	// Rutas públicas
	r.HandleFunc("/api/login", handlers.Login).Methods("POST")

	// Rutas de Students (públicas para el frontend)
	r.HandleFunc("/api/students", studentHandler.GetAllStudents).Methods("GET")
	r.HandleFunc("/api/students", studentHandler.CreateStudent).Methods("POST")
	r.HandleFunc("/api/students/{id}", studentHandler.GetStudent).Methods("GET")
	r.HandleFunc("/api/students/{id}", studentHandler.UpdateStudent).Methods("PUT")
	r.HandleFunc("/api/students/{id}", studentHandler.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/api/students/{id}/confirm", studentHandler.ConfirmStudent).Methods("POST")

	// Rutas de Groups (públicas para el frontend)
	r.HandleFunc("/api/groups", groupHandler.GetAllGroups).Methods("GET")
	r.HandleFunc("/api/groups", groupHandler.CreateGroup).Methods("POST")
	r.HandleFunc("/api/groups/{id}", groupHandler.GetGroup).Methods("GET")
	r.HandleFunc("/api/groups/{id}", groupHandler.UpdateGroup).Methods("PUT")
	r.HandleFunc("/api/groups/{id}", groupHandler.DeleteGroup).Methods("DELETE")

	// Rutas de Invitations (públicas para el frontend)
	r.HandleFunc("/api/invitations", invitationHandler.GetAllInvitations).Methods("GET")
	r.HandleFunc("/api/invitations", invitationHandler.CreateInvitation).Methods("POST")
	r.HandleFunc("/api/invitations/{groupId}", invitationHandler.UpdateInvitation).Methods("PUT")

	// Rutas protegidas de administración
	admin := r.PathPrefix("/api/admin").Subrouter()
	admin.Use(handlers.AuthMiddleware)

	admin.HandleFunc("/students", studentHandler.GetAllStudents).Methods("GET")
	admin.HandleFunc("/students", studentHandler.CreateStudent).Methods("POST")
	admin.HandleFunc("/students/{id}", studentHandler.UpdateStudent).Methods("PUT")
	admin.HandleFunc("/students/{id}", studentHandler.DeleteStudent).Methods("DELETE")
	admin.HandleFunc("/students/import", studentHandler.ImportStudents).Methods("POST")

	admin.HandleFunc("/groups", groupHandler.GetAllGroups).Methods("GET")
	admin.HandleFunc("/groups/{id}", groupHandler.DeleteGroup).Methods("DELETE")
	admin.HandleFunc("/groups/{id}/members/{memberId}", groupHandler.RemoveMember).Methods("DELETE")

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
