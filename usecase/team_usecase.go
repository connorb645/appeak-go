package usecase

import (
	"context"
	"time"

	"github.com/connorb645/appeak-go/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type teamUsecase struct {
	teamRepository domain.TeamRepository
	contextTimeout time.Duration
}

func NewTeamUsecase(tr domain.TeamRepository, timeout time.Duration) domain.TeamUsecase {
	return &teamUsecase{
		teamRepository: tr,
		contextTimeout: timeout,
	}
}

func (tu *teamUsecase) Create(c context.Context, team *domain.Team) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.teamRepository.Create(ctx, team)
}

func (tu *teamUsecase) Fetch(c context.Context, teamID primitive.ObjectID) (*domain.Team, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.teamRepository.Fetch(ctx, teamID)
}