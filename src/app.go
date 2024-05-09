package src

import (
	"fmt"
	"log"
	"net/http"
	"restapi-altera/config"
	"restapi-altera/src/delivery"
	"restapi-altera/src/helper"
	"restapi-altera/src/helper/validator"
	"restapi-altera/src/repository"
	"restapi-altera/src/usecase"

	validatorEngine "github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server interface {
	Run()
}

type server struct {
	httpServer *echo.Echo
	config     config.Config
}

func InitServer(config config.Config) Server {

	e := echo.New()
	e.Validator = &validator.GoPlaygroundValidator{
		Validator: validatorEngine.New(),
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &server{
		config:     config,
		httpServer: e,
	}

}

func (s *server) Run() {

	//tester route
	s.httpServer.GET("/", func(c echo.Context) error {
		return helper.WriteResponse(c, http.StatusOK, "hallo dari echo", nil)
	})

	repo := repository.NewRepository(s.config)
	usecase := usecase.NewUsecase(repo)
	delivery := delivery.NewDelivery(usecase)
	delivery.Auth(s.httpServer.Group("auth"))
	delivery.User(s.httpServer.Group("user"))

	if err := s.httpServer.Start(fmt.Sprintf(":%d", s.config.Port())); err != nil {
		log.Panic(err)
	}
}
