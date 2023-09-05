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
	setupPublicRoutes(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	setupProtectedRoutes(env, timeout, db, store, protectedRouter)

	protectedTeamAdminRouter := protectedRouter.Group("")
	setupTeamAdminRoutes(env, timeout, db, protectedTeamAdminRouter)
}

func setupPublicRoutes(env *bootstrap.Env, timeout time.Duration, db mongo.Database, router *gin.RouterGroup) {
	// All Public APIs
	NewSignupRouter(env, timeout, db, router)
	NewLoginRouter(env, timeout, db, router)
	NewRefreshTokenRouter(env, timeout, db, router)
}

func setupProtectedRoutes(env *bootstrap.Env, timeout time.Duration, db mongo.Database, store store.HelpCenterProvider, router *gin.RouterGroup) {
	// Middleware to verify AccessToken
	router.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(timeout, db, router)
	NewTaskRouter(timeout, db, router)
	NewDocumentRouter(timeout, store, router)
	NewTeamCreationRouter(env, timeout, db, router)
}

func setupTeamAdminRoutes(env *bootstrap.Env, timeout time.Duration, db mongo.Database, router *gin.RouterGroup) {
	router.Use(middleware.TeamAdminMiddleware(db))
	NewTeamManagementRouter(env, timeout, db, router)
}
