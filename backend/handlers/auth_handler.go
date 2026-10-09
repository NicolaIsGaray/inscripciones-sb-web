package handlers

import (
	"encoding/json"
	"net/http"

	"servidor-angular/config"
)

// LoginRequest representa la solicitud de login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse representa la respuesta del login
type LoginResponse struct {
	Token string `json:"token"`
}

// Login maneja el inicio de sesión del administrador
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
		return
	}

	// Obtener credenciales del administrador
	creds, err := config.GetAdminCredentials()
	if err != nil {
		http.Error(w, "Error de configuración del servidor", http.StatusInternalServerError)
		return
	}

	// Verificar credenciales
	if req.Username != creds.Username || !config.VerifyPassword(req.Password, creds.Password) {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	// Generar token
	token, err := config.GenerateToken(req.Username)
	if err != nil {
		http.Error(w, "Error al generar token", http.StatusInternalServerError)
		return
	}

	// Responder con el token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}

// AuthMiddleware verifica el token JWT en las rutas protegidas
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No autorizado", http.StatusUnauthorized)
			return
		}

		// Extraer token del header "Bearer <token>"
		tokenString := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}

		if tokenString == "" {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		// Validar token
		_, err := config.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Token inválido o expirado", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
