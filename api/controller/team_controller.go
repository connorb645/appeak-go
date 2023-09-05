package controller

import (
	"context"
	"net/http"
	"sync"

	"github.com/connorb645/appeak-go/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileGetter interface {
	GetProfileByID(c context.Context, userID primitive.ObjectID) (*domain.Profile, error)
}

type TeamController struct {
	TeamUsecase domain.TeamUsecase
	ProfileGetter ProfileGetter
}

func (tc *TeamController) Create(c *gin.Context) {
	var team domain.Team

	err := c.ShouldBind(&team)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	userID := c.GetString("x-user-id")
	team.ID = primitive.NewObjectID()
	
	team.AdminID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: userID,})
		return
	}
	team.TeamMembers = []primitive.ObjectID{team.AdminID}

	err = tc.TeamUsecase.Create(c, &team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Team created successfully",
	})
}

func (tc *TeamController) Fetch(c *gin.Context) {
	teamID, err := primitive.ObjectIDFromHex(c.Param("teamID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	team, err := tc.TeamUsecase.Fetch(c, teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}

func (tc *TeamController) Update(c *gin.Context) {
	var team domain.TeamUpdate

	err := c.ShouldBind(&team)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	teamID, err := primitive.ObjectIDFromHex(c.Param("teamID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	validTeamMembers := tc.fetchValidTeamMembers(c, team.TeamMembers)
	team.TeamMembers = &validTeamMembers

	err = tc.TeamUsecase.Update(c, teamID, &team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Team updated successfully",
	})
}

func (tc *TeamController) fetchValidTeamMembers(c *gin.Context, teamMembers *[]primitive.ObjectID) []primitive.ObjectID {
	resultCh := make(chan primitive.ObjectID, len(*teamMembers))
	var wg sync.WaitGroup

	for _, memberID := range *teamMembers {
		wg.Add(1)
		go func(memberID primitive.ObjectID) {
			defer wg.Done()

			p, err := tc.ProfileGetter.GetProfileByID(c, memberID)
			if err == nil && p != nil {
				resultCh <- memberID
			}
		}(memberID)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	validTeamMembers := make([]primitive.ObjectID, 0)
	for memberID := range resultCh {
		validTeamMembers = append(validTeamMembers, memberID)
	}

	return validTeamMembers
}
