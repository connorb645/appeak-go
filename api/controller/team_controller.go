package controller

import (
	"net/http"

	"github.com/connorb645/appeak-go/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamController struct {
	TeamUsecase domain.TeamUsecase
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
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
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