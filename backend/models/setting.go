package models

// Settings representa la configuración global de la aplicación.
// Es un documento único en la colección "settings" con _id fijo "app-settings".
type Settings struct {
	ID            string `json:"id" bson:"_id"`
	GroupsEnabled bool   `json:"groupsEnabled" bson:"groupsEnabled"`
}