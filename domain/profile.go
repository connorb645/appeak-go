package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID primitive.ObjectID) (*Profile, error)
}
