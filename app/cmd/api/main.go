package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wasuwa/blog-api/app/config"
	"github.com/wasuwa/blog-api/app/infra/persistence"
	"github.com/wasuwa/blog-api/app/interfaces/handler"
	"github.com/wasuwa/blog-api/app/redis"
	"github.com/wasuwa/blog-api/app/usecase"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	client := redis.NewClient()
	postPersistence := persistence.NewPostPersistence(client)
	postUseCase := usecase.NewPostUseCase(postPersistence)
	postHandler := handler.NewPostHandler(postUseCase)
	postHandler.Routes(e)

	e.Logger.Fatal(e.Start(config.APIPort))
}
