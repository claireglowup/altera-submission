package delivery

import (
	"restapi-altera/src/usecase"

	"github.com/labstack/echo/v4"
)

type Delivery interface {
	Auth(e *echo.Group)
	User(e *echo.Group)
}

type delivery struct {
	usecase usecase.Usecase
}

func NewDelivery(usecase usecase.Usecase) Delivery {
	return &delivery{
		usecase: usecase,
	}
}
