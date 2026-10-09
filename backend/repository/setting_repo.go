package repository

import (
	"context"
	"time"

	"servidor-angular/config"
	"servidor-angular/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// settingsID es el _id fijo del documento único de configuración
const settingsID = "app-settings"

// SettingRepository handles database operations for the settings document
type SettingRepository struct {
	collection *mongo.Collection
}

// NewSettingRepository creates a new setting repository
func NewSettingRepository() *SettingRepository {
	return &SettingRepository{
		collection: config.GetCollection("settings"),
	}
}

// Get returns the settings document. Si todavía no existe lo crea con
// los valores por defecto (grupos deshabilitados).
func (r *SettingRepository) Get() (*models.Settings, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var settings models.Settings
	err := r.collection.FindOne(ctx, bson.M{"_id": settingsID}).Decode(&settings)

	if err == mongo.ErrNoDocuments {
		settings = models.Settings{ID: settingsID, GroupsEnabled: false}
		if _, err := r.collection.InsertOne(ctx, settings); err != nil {
			return nil, err
		}
		return &settings, nil
	}

	if err != nil {
		return nil, err
	}

	return &settings, nil
}

// SetGroupsEnabled actualiza el flag de habilitación de grupos
func (r *SettingRepository) SetGroupsEnabled(enabled bool) (*models.Settings, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	settings := models.Settings{ID: settingsID, GroupsEnabled: enabled}
	_, err := r.collection.ReplaceOne(
		ctx,
		bson.M{"_id": settingsID},
		settings,
		options.Replace().SetUpsert(true),
	)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}