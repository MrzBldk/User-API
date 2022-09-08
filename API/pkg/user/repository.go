package user

import (
	"context"
	"time"

	"github.com/MrzBldk/User-API/api/presenter"
	"github.com/MrzBldk/User-API/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	ReadUser(id string) (*presenter.User, error)
	ReadUsers() (*[]presenter.User, error)
	UpdateUser(user *entities.User) error
	DeleteUser(id string) error
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreateUser(user *entities.User) (*entities.User, error) {
	user.Id = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) ReadUser(id string) (*presenter.User, error) {
	var user presenter.User
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.Collection.FindOne(context.Background(), bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) ReadUsers() (*[]presenter.User, error) {
	var users []presenter.User
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user presenter.User
		_ = cursor.Decode(&user)
		users = append(users, user)
	}
	return &users, nil
}

func (r *repository) UpdateUser(user *entities.User) error {
	user.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": user.Id}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteUser(id string) error {
	UserId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": UserId})
	if err != nil {
		return err
	}
	return nil
}
