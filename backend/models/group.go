package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Group representa un grupo de estudiantes
type Group struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Leader  Student            `json:"leader" bson:"leader"`
	Members []Student          `json:"members,omitempty" bson:"members,omitempty"`
	Pending []Student          `json:"pending,omitempty" bson:"pending,omitempty"`
	Full    bool               `json:"full,omitempty" bson:"full,omitempty"`
}
