package router

import (
	"net/http"
	"os"
	"strings"

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

	// Rutas de Students: SOLO lectura y confirmación son públicas.
	// Crear, editar y borrar viven bajo /api/admin (protegidas con JWT).
	r.HandleFunc("/api/students", studentHandler.GetAllStudents).Methods("GET")
	r.HandleFunc("/api/students/{id}", studentHandler.GetStudent).Methods("GET")
	r.HandleFunc("/api/students/{id}/confirm", studentHandler.ConfirmStudent).Methods("POST")

	// Rutas de Groups: SOLO lectura es pública
	r.HandleFunc("/api/groups", groupHandler.GetAllGroups).Methods("GET")
	r.HandleFunc("/api/groups/{id}", groupHandler.GetGroup).Methods("GET")

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
	admin.HandleFunc("/groups", groupHandler.CreateGroup).Methods("POST")
	admin.HandleFunc("/groups/{id}", groupHandler.UpdateGroup).Methods("PUT")
	admin.HandleFunc("/groups/{id}", groupHandler.DeleteGroup).Methods("DELETE")
	admin.HandleFunc("/groups/{id}/members/{memberId}", groupHandler.RemoveMember).Methods("DELETE")

	return r
}

// allowedOrigins devuelve los orígenes permitidos, leídos de ALLOWED_ORIGINS.
// Si la variable no está definida devuelve nil, que significa "cualquiera"
// (solo aceptable en desarrollo).
func allowedOrigins() []string {
	raw := os.Getenv("ALLOWED_ORIGINS")
	if strings.TrimSpace(raw) == "" {
		return nil
	}

	var origins []string
	for _, origin := range strings.Split(raw, ",") {
		if trimmed := strings.TrimSpace(origin); trimmed != "" {
			origins = append(origins, trimmed)
		}
	}
	return origins
}

// normalizeOrigin deja un origen comparable: minúsculas y sin esquema.
// Permite configurar ALLOWED_ORIGINS como "https://sitio.onrender.com" o
// simplemente "sitio.onrender.com" (que es como Render expone el hostname).
func normalizeOrigin(value string) string {
	value = strings.TrimSpace(strings.ToLower(value))
	for _, scheme := range []string{"https://", "http://"} {
		if strings.HasPrefix(value, scheme) {
			return strings.TrimSuffix(value[len(scheme):], "/")
		}
	}
	return strings.TrimSuffix(value, "/")
}

// isOriginAllowed dice si el origen de la petición está permitido.
// Sin lista configurada, se permite cualquier origen.
func isOriginAllowed(origin string, allowed []string) bool {
	if len(allowed) == 0 {
		return true
	}
	normalized := normalizeOrigin(origin)
	for _, candidate := range allowed {
		if normalizeOrigin(candidate) == normalized {
			return true
		}
	}
	return false
}

// CorsMiddleware handles CORS headers
func CorsMiddleware(next http.Handler) http.Handler {
	allowed := allowedOrigins()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Si el origen no está permitido, no se le devuelve ninguna cabecera CORS:
		// el navegador bloqueará la respuesta.
		if origin != "" && !isOriginAllowed(origin, allowed) {
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			http.Error(w, "Origen no permitido", http.StatusForbidden)
			return
		}

		if len(allowed) == 0 {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		} else if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Add("Vary", "Origin")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
