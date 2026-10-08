# Variables
FRONTEND_DIR = frontend
BACKEND_DIR = backend
BACKEND_PORT = 8080
FRONTEND_PORT = 4200

# Comandos principales
.PHONY: all backend frontend dev build-backend build-frontend clean

all: backend frontend

# Levantar backend
backend:
	@echo "Levantando backend en puerto $(BACKEND_PORT)..."
	cd $(BACKEND_DIR) && go run main.go

# Levantar frontend
frontend:
	@echo "Levantando frontend en puerto $(FRONTEND_PORT)..."
	cd $(FRONTEND_DIR) && ng serve

# Levantar ambos en paralelo
dev:
	@echo "Levantando frontend y backend..."
	@make backend & make frontend & wait

# Compilar backend
build-backend:
	@echo "Compilando backend..."
	cd $(BACKEND_DIR) && go build -o ../bin/server main.go

# Compilar frontend
build-frontend:
	@echo "Compilando frontend..."
	cd $(FRONTEND_DIR) && ng build

# Limpiar
clean:
	rm -rf bin/
	rm -rf $(FRONTEND_DIR)/dist/
