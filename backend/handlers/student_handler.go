package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"strings"

	"servidor-angular/models"
	"servidor-angular/repository"

	"github.com/gorilla/mux"
	"github.com/xuri/excelize/v2"
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

// normalizeSchool normaliza el valor de escuela a "secundaria" o "deportiva".
// Devuelve "" si el valor no corresponde a ninguna escuela válida.
func normalizeSchool(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "secundaria", "escuela secundaria", "secondary":
		return "secundaria"
	case "deportiva", "escuela deportiva", "deporte", "sports":
		return "deportiva"
	default:
		return ""
	}
}

// isEmptyRow indica si todas las celdas de la fila están vacías
func isEmptyRow(row []string) bool {
	for _, cell := range row {
		if strings.TrimSpace(cell) != "" {
			return false
		}
	}
	return true
}

// containsSheet indica si el nombre de hoja está en la lista
func containsSheet(sheets []string, name string) bool {
	for _, sheet := range sheets {
		if sheet == name {
			return true
		}
	}
	return false
}

// GetAllStudents handles GET /api/students</path>
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

	// Normalizar escuela (por defecto "secundaria" si no viene informada)
	if s := normalizeSchool(student.School); s != "" {
		student.School = s
	} else {
		student.School = "secundaria"
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

	// Normalizar escuela (por defecto "secundaria" si no viene informada)
	if s := normalizeSchool(student.School); s != "" {
		student.School = s
	} else {
		student.School = "secundaria"
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

// ConfirmStudent handles POST /api/students/{id}/confirm
// Marca el banco como confirmado y guarda el email que ingresa el estudiante.
// Es público porque los ingresantes no tienen login: por eso solo permite
// modificar `confirmed` y `email`, nunca el resto de la ficha.
func (h *StudentHandler) ConfirmStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var payload struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Error en los datos enviados", http.StatusBadRequest)
		return
	}

	email := strings.TrimSpace(payload.Email)
	if email == "" {
		http.Error(w, "El correo electrónico es obligatorio", http.StatusBadRequest)
		return
	}

	if _, err := mail.ParseAddress(email); err != nil {
		http.Error(w, "Correo electrónico inválido", http.StatusBadRequest)
		return
	}

	student, err := h.repo.Confirm(id, email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Error al confirmar el banco", http.StatusInternalServerError)
		return
	}

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

// ImportStudents handles POST /api/admin/students/import
func (h *StudentHandler) ImportStudents(w http.ResponseWriter, r *http.Request) {
	// Parsear el archivo multipart
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Error al leer el archivo", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error al obtener el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Verificar extensión
	if !strings.HasSuffix(handler.Filename, ".xlsx") {
		http.Error(w, "Solo se permiten archivos .xlsx", http.StatusBadRequest)
		return
	}

	// Abrir el archivo Excel
	f, err := excelize.OpenReader(file)
	if err != nil {
		http.Error(w, "Error al abrir el archivo Excel", http.StatusBadRequest)
		return
	}
	defer f.Close()

	// Buscar la hoja de datos: preferimos "Hoja1" y, si no existe, usamos la primera
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		http.Error(w, "El archivo no tiene ninguna hoja", http.StatusBadRequest)
		return
	}

	sheet := "Hoja1"
	if !containsSheet(sheets, sheet) {
		sheet = sheets[0]
	}

	rows, err := f.GetRows(sheet)
	if err != nil {
		http.Error(w, "Error al leer las filas del Excel", http.StatusBadRequest)
		return
	}

	// Una plantilla vacía (solo encabezados) es válida: simplemente no importa alumnos
	if len(rows) < 1 {
		http.Error(w, "El archivo no tiene filas", http.StatusBadRequest)
		return
	}

	// La escuela solo se toma de la columna E del archivo. Si la celda está
	// vacía o trae un valor no reconocido, el alumno se importa como secundaria.
	const defaultSchool = "secundaria"

	// DNI ya registrados, para no duplicar alumnos al reimportar un archivo
	existingDNIs, err := h.repo.FindAllDNIs()
	if err != nil {
		http.Error(w, "Error al verificar alumnos existentes", http.StatusInternalServerError)
		return
	}
	seenDNIs := make(map[string]bool, len(existingDNIs))
	for _, dni := range existingDNIs {
		seenDNIs[dni] = true
	}

	// Campos obligatorios: dni, name, instance, title
	var successCount int
	var errorMessages []string

	for i, row := range rows {
		if i == 0 {
			continue // Saltar encabezados
		}

		// Filas totalmente vacías (separadores) no son un error
		if isEmptyRow(row) {
			continue
		}

		// Solo A-D son obligatorias: la columna E (escuela) es opcional y
		// Excel recorta la fila cuando su última celda queda vacía.
		if len(row) < 4 {
			errorMessages = append(errorMessages, fmt.Sprintf("Fila %d: Faltan columnas", i+1))
			continue
		}

		dni := strings.TrimSpace(row[0])
		name := strings.TrimSpace(row[1])
		instance := strings.TrimSpace(row[2])
		title := strings.TrimSpace(row[3])

		// Validar campos obligatorios
		if dni == "" || name == "" || instance == "" || title == "" {
			errorMessages = append(errorMessages, fmt.Sprintf("Fila %d: Faltan campos obligatorios", i+1))
			continue
		}

		// El DNI identifica al ingresante: si ya existe (en la base o en otra fila
		// del mismo archivo) no lo duplicamos.
		if seenDNIs[dni] {
			errorMessages = append(errorMessages, fmt.Sprintf("Fila %d: El DNI %s ya está cargado, se omitió", i+1, dni))
			continue
		}
		seenDNIs[dni] = true

		// La escuela viene de la columna E; si no es válida, queda secundaria.
		// El ingresante aún no verificó su email ni confirmó el banco, por eso
		// ambos campos empiezan vacíos / en false.
		school := defaultSchool
		if len(row) > 4 {
			if s := normalizeSchool(strings.TrimSpace(row[4])); s != "" {
				school = s
			}
		}

		student := models.Student{
			DNI:       dni,
			Name:      name,
			Email:     "",
			Instance:  instance,
			Title:     title,
			School:    school,
			Confirmed: false,
		}

		if err := h.repo.Create(&student); err != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("Fila %d: Error al crear alumno: %v", i+1, err))
			continue
		}

		successCount++
	}

	// Los errores siempre travelan como lista, nunca null
	if errorMessages == nil {
		errorMessages = []string{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": successCount,
		"errors":  errorMessages,
	})
}
