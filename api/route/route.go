package route

import (
	"time"

	"github.com/connorb645/appeak-go/api/middleware"
	"github.com/connorb645/appeak-go/bootstrap"
	"github.com/connorb645/appeak-go/mongo"
	"github.com/connorb645/appeak-go/store"
	"github.com/gin-gonic/gin"
)

func Setup(
	env *bootstrap.Env, 
	timeout time.Duration, 
	db mongo.Database, 
	store store.HelpCenterProvider, 
	gin *gin.Engine,
) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(timeout, db, protectedRouter)
	NewTaskRouter(timeout, db, protectedRouter)
	NewDocumentRouter(timeout, store, protectedRouter)
}
