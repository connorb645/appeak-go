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
	Update(c context.Context, teamID primitive.ObjectID, team *TeamUpdate) error
}

type TeamUsecase interface {
	Create(c context.Context, team *Team) error
	Fetch(c context.Context, teamID primitive.ObjectID) (*Team, error)
	Update(c context.Context, teamID primitive.ObjectID, team *TeamUpdate) error
}

type Team struct {
	ID      primitive.ObjectID 		 `bson:"_id" json:"-"`
	Title   string             		 `bson:"title" json:"title" binding:"required"`
	AdminID primitive.ObjectID 		 `bson:"adminID" json:"adminID"`
	TeamMembers []primitive.ObjectID `bson:"teamMembers" json:"teamMembers"`
}

type TeamUpdate struct {
	Title       *string                  `bson:"title,omitempty" json:"title,omitempty"`
	AdminID     *primitive.ObjectID 	 `bson:"adminID,omitempty" json:"adminID,omitempty"`
	TeamMembers *[]primitive.ObjectID    `bson:"teamMembers,omitempty" json:"teamMembers,omitempty"`
}