package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Student representa un estudiante en el sistema
type Student struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DNI         string             `json:"dni" bson:"dni"`
	Name        string             `json:"name" bson:"name"`
	Email       string             `json:"email" bson:"email"`
	Instance    string             `json:"instance" bson:"instance"`
	Title       string             `json:"title" bson:"title"`
	School      string             `json:"school" bson:"school"`
	HasGroup    bool               `json:"hasGroup" bson:"hasGroup"`
	Alone       bool               `json:"alone" bson:"alone"`
	Confirmed   bool               `json:"confirmed" bson:"confirmed"`
	Invitations []Invitation       `json:"invitations,omitempty" bson:"invitations,omitempty"`
}
