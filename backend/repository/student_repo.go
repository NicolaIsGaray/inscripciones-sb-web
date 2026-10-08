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

// StudentRepository handles database operations for students
type StudentRepository struct {
	collection *mongo.Collection
}

// NewStudentRepository creates a new student repository
func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		collection: config.GetCollection("students"),
	}
}

// FindAll returns all students
func (r *StudentRepository) FindAll() ([]models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var students []models.Student
	if err = cursor.All(ctx, &students); err != nil {
		return nil, err
	}

	if students == nil {
		students = []models.Student{}
	}

	return students, nil
}

// FindByID returns a student by ID
func (r *StudentRepository) FindByID(id primitive.ObjectID) (*models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var student models.Student
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

// Create inserts a new student
func (r *StudentRepository) Create(student *models.Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.InsertOne(ctx, student)
	if err != nil {
		return err
	}

	student.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

// Update modifies an existing student
func (r *StudentRepository) Update(id primitive.ObjectID, student *models.Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"dni":       student.DNI,
			"name":      student.Name,
			"email":     student.Email,
			"instance":  student.Instance,
			"title":     student.Title,
			"hasGroup":  student.HasGroup,
			"alone":     student.Alone,
			"confirmed": student.Confirmed,
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

// Delete removes a student
func (r *StudentRepository) Delete(id primitive.ObjectID) error {
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

// GetCollection returns the MongoDB collection for advanced operations
func (r *StudentRepository) GetCollection() *mongo.Collection {
	return r.collection
}
