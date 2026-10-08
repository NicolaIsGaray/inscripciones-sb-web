package repository

import (
	"context"
	"time"

	"servidor-angular/config"
	"servidor-angular/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GroupRepository handles database operations for groups
type GroupRepository struct {
	collection *mongo.Collection
}

// NewGroupRepository creates a new group repository
func NewGroupRepository() *GroupRepository {
	return &GroupRepository{
		collection: config.GetCollection("groups"),
	}
}

// FindAll returns all groups
func (r *GroupRepository) FindAll() ([]models.Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var groups []models.Group
	if err = cursor.All(ctx, &groups); err != nil {
		return nil, err
	}

	if groups == nil {
		groups = []models.Group{}
	}

	return groups, nil
}

// FindByID returns a group by ID
func (r *GroupRepository) FindByID(id primitive.ObjectID) (*models.Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var group models.Group
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

// Create inserts a new group
func (r *GroupRepository) Create(group *models.Group) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, group)
	if err != nil {
		return err
	}

	group.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

// Update modifies an existing group
func (r *GroupRepository) Update(id primitive.ObjectID, group *models.Group) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"leader":  group.Leader,
			"members": group.Members,
			"pending": group.Pending,
			"full":    group.Full,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

// Delete removes a group
func (r *GroupRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
