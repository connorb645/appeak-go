package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTeam = "teams"
)

type TeamRepository interface {
	Create(c context.Context, team *Team) error
}

type TeamUsecase interface {
	Create(c context.Context, team *Team) error
}

type Team struct {
	ID      primitive.ObjectID `bson:"_id" json:"-"`
	Title   string             `bson:"title" form:"title" binding:"required" json:"title"`
	AdminID primitive.ObjectID `bson:"adminID" json:"-"`
}