package validator

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type GoPlaygroundValidator struct {
	Validator *validator.Validate
}

func (gp *GoPlaygroundValidator) Validate(i interface{}) error {
	if err := gp.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
