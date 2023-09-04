package route

import (
	"time"

	"github.com/connorb645/appeak-go/api/controller"
	"github.com/connorb645/appeak-go/bootstrap"
	"github.com/connorb645/appeak-go/domain"
	"github.com/connorb645/appeak-go/mongo"
	"github.com/connorb645/appeak-go/repository"
	"github.com/connorb645/appeak-go/usecase"
	"github.com/gin-gonic/gin"
)

func NewTeamRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTeamRepository(db, domain.CollectionTeam)
	tc := &controller.TeamController{
		TeamUsecase: usecase.NewTeamUsecase(tr, timeout),
	}
	group.POST("/teams", tc.Create)
}