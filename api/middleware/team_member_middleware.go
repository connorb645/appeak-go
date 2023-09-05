package middleware

import (
	"net/http"

	"github.com/connorb645/appeak-go/domain"
	"github.com/connorb645/appeak-go/mongo"
	"github.com/connorb645/appeak-go/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TeamAdminMiddleware(db mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		teamID, err := primitive.ObjectIDFromHex(c.Param("teamID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
			c.Abort()
			return
		}

		userID, _ :=primitive.ObjectIDFromHex(c.GetString("x-user-id")) 

		tr := repository.NewTeamRepository(db, domain.CollectionTeam)
		team, err := tr.Fetch(c, teamID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
			c.Abort()
			return
		}

		if team.AdminID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": "User is not the team admin"})
			c.Abort()
			return
		}

		c.Next()
	}
}
