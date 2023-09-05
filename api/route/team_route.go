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

func makeTeamController(db mongo.Database, timeout time.Duration) *controller.TeamController {
	tr := repository.NewTeamRepository(db, domain.CollectionTeam)
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	tc := &controller.TeamController{
		TeamUsecase:   usecase.NewTeamUsecase(tr, timeout),
		ProfileGetter: usecase.NewProfileUsecase(ur, timeout),
	}
	return tc
}

func NewTeamCreationRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tc := makeTeamController(db, timeout)
	group.POST("/teams", tc.Create)
}

func NewTeamManagementRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tc := makeTeamController(db, timeout)
	group.GET("/teams/:teamID", tc.Fetch)
	group.PATCH("/teams/:teamID", tc.Update)
}