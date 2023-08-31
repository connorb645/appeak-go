package main

import (
	"time"

	route "github.com/connorb645/appeak-go/api/route"
	"github.com/connorb645/appeak-go/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	defer app.CloseDBConnection()

	env := app.Env

	db := app.Mongo.Database(env.DBName)

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)
	
	gin.Run(env.ServerAddress)
}
