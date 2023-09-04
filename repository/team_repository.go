package repository

import (
	"context"

	"github.com/connorb645/appeak-go/domain"
	"github.com/connorb645/appeak-go/mongo"
)

type teamRepository struct {
	database   mongo.Database
	collection string
}

func NewTeamRepository(db mongo.Database, collection string) domain.TeamRepository {
	return &teamRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *teamRepository) Create(c context.Context, team *domain.Team) error {
	collection := tr.database.Collection(tr.collection)
	_, err := collection.InsertOne(c, team)
	return err
}