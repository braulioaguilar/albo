package main

import (
	"albo/colaborator"
	"albo/colaborator/handler"
	"albo/colaborator/repository"
	"albo/colaborator/service"
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

func NewApplication() (*Application, error) {
	db, _ := mongo.Connect()
	repo := repository.NewRepository(db, context.Background())

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
	app, err := NewApplication()
	if err != nil {
		log.Println(err)
	}

	// Sync with interval time, see ENV
	go sync.Marvel(app.Service)

	if err := app.Run(); err != nil {
		log.Panicln(err)
	}
}
