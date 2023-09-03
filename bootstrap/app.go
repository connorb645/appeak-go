package bootstrap

import (
	"github.com/connorb645/appeak-go/mongo"
	"github.com/connorb645/appeak-go/store"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
	Store store.HelpCenterProvider
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	app.Store = NewDocumentStore()
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
