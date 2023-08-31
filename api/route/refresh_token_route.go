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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
