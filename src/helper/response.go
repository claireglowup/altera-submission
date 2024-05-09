package helper

import (
	"restapi-altera/src/model/response"

	"github.com/labstack/echo/v4"
)

func WriteResponse(c echo.Context, status int, message interface{}, data interface{}) error {
	return c.JSON(status, &response.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
