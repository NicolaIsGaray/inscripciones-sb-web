package config

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB holds the database connection
var DB *mongo.Database

// Connect establishes connection to MongoDB Atlas
func Connect() error {
	// Obtener URI desde variable de entorno
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return errors.New("la variable de entorno MONGODB_URI no está definida")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conectar a MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// Verificar conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	// Obtener referencia a la base de datos
	DB = client.Database("inscripciones")

	log.Println("Conexión a MongoDB Atlas establecida correctamente")
	return nil
}

// GetCollection returns a MongoDB collection
func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}

// Disconnect closes the MongoDB connection
func Disconnect() error {
	if DB != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return DB.Client().Disconnect(ctx)
	}
	return nil
}
