package repository

import (
	"context"

	"github.com/connorb645/appeak-go/domain"
	"github.com/connorb645/appeak-go/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamRepository struct {
	Database   mongo.Database
	Collection string
}

func NewTeamRepository(db mongo.Database, collection string) domain.TeamRepository {
	return &TeamRepository{
		Database:   db,
		Collection: collection,
	}
}

func (tr *TeamRepository) Create(c context.Context, team *domain.Team) error {
	collection := tr.Database.Collection(tr.Collection)
	_, err := collection.InsertOne(c, team)
	return err
}

func (tr *TeamRepository) Fetch(
	c context.Context, 
	teamID primitive.ObjectID,
) (*domain.Team, error) {
	collection := tr.Database.Collection(tr.Collection)
	sr := collection.FindOne(c, bson.M{"_id": teamID})
	var team domain.Team
	err := sr.Decode(&team)
	if err != nil {
		return nil, err
	}
	return &team, nil
}