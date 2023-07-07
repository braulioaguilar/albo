package main

import (
	"albo/colaborator"
	"albo/colaborator/handler"
	"albo/colaborator/repository"
	"albo/colaborator/service"
	c "albo/config"
	"albo/pkg/mongo"
	"albo/sync"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Router  *gin.Engine
	Service colaborator.Service
}

func NewApplication(ctx context.Context) (*Application, error) {
	db, err := mongo.Connect(ctx, c.Config.MONGO_URI)
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(db.Database(c.Config.DB_NAME), ctx)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return &Application{
		Router:  router,
		Service: service.NewColaboratorService(repo),
	}, nil
}

func (app *Application) Run() error {
	handler.MakeHandler(app.Service, app.Router)
	return app.Router.Run(":8080")
}

func main() {
	ctx := context.Background()
	app, err := NewApplication(ctx)
	if err != nil {
		log.Println(err)
	}

	// Sync with interval time, see ENV
	go sync.Marvel(app.Service)

	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
