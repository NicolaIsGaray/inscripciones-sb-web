package models

// Invitation representa una invitación a un grupo
type Invitation struct {
	GroupID    string   `json:"groupId" bson:"groupId"`
	LeaderName string   `json:"leaderName" bson:"leaderName"`
	Members    []string `json:"members" bson:"members"`
	Status     string   `json:"status,omitempty" bson:"status,omitempty"`
}
