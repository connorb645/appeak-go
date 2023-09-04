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
	Fetch(c context.Context, teamID primitive.ObjectID) (*Team, error)
}

type TeamUsecase interface {
	Create(c context.Context, team *Team) error
	Fetch(c context.Context, teamID primitive.ObjectID) (*Team, error)
}

type Team struct {
	ID      primitive.ObjectID 		 `bson:"_id" json:"-"`
	Title   string             		 `bson:"title" json:"title" binding:"required"`
	AdminID primitive.ObjectID 		 `bson:"adminID" json:"adminID"`
	TeamMembers []primitive.ObjectID `bson:"teamMembers" json:"teamMembers"`
}