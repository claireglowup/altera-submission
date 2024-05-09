package delivery

import (
	"fmt"
	"net/http"
	"restapi-altera/src/helper"
	"restapi-altera/src/model/request"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func (d *delivery) GetUserByID(c echo.Context) error {

	ctx := c.Request().Context()

	cookie, err := c.Cookie("jwt")
	if err != nil {
		return helper.WriteResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", nil)
	}

	data, err := d.usecase.GetUserByID(ctx, cookie.Value)
	if err != nil {
		return helper.WriteResponse(c, http.StatusNotFound, err.Error(), nil)
	}

	return helper.WriteResponse(c, http.StatusOK, fmt.Sprintf("hi %s", data.Name), data)

}

func (d *delivery) GetUserAllShoes(c echo.Context) error {

	ctx := c.Request().Context()

	cookie, err := c.Cookie("jwt")
	if err != nil {
		return helper.WriteResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", nil)
	}

	data, err := d.usecase.GetUserAllShoes(ctx, cookie.Value)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return helper.WriteResponse(c, http.StatusOK, fmt.Sprintf("hi %s", data.Name), data)

}

func (d *delivery) GetUserShoesByNoId(c echo.Context) error {

	ctx := c.Request().Context()

	cookie, err := c.Cookie("jwt")
	if err != nil {
		return helper.WriteResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", nil)
	}

	noId, _ := strconv.Atoi(c.Param("no_id"))

	data, err := d.usecase.GetUserShoesByNoId(ctx, noId, cookie.Value)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return helper.WriteResponse(c, http.StatusNotFound, fmt.Sprintf("shoes no_id = %d is not found", noId), nil)
		}
		return helper.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)

	}

	return helper.WriteResponse(c, http.StatusOK, fmt.Sprintf("hi %s", data.Name), data)

}

func (d *delivery) AddNewShoes(c echo.Context) error {

	var payload *request.ReqUserShoes

	ctx := c.Request().Context()

	cookie, err := c.Cookie("jwt")
	if err != nil {
		return helper.WriteResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", nil)
	}

	err = c.Bind(&payload)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err = c.Validate(payload)
	if err != nil {
		return helper.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	err = d.usecase.AddUserShoes(ctx, cookie.Value, payload)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return helper.WriteResponse(c, http.StatusOK, "new shoes successfully added", nil)

}

func (d *delivery) UpdateShoesByNoId(c echo.Context) error {

	var payload *request.UpdateUserShoes

	ctx := c.Request().Context()

	noId, _ := strconv.Atoi(c.Param("no_id"))

	cookie, err := c.Cookie("jwt")
	if err != nil {
		return helper.WriteResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", nil)
	}

	err = c.Bind(&payload)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err = c.Validate(payload)
	if err != nil {
		return helper.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	err = d.usecase.UpdateShoesByNoId(ctx, noId, cookie.Value, payload)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return helper.WriteResponse(c, http.StatusOK, "succesfully updated data", nil)

}

func (d *delivery) DeleteShoesByNoId(c echo.Context) error {

	ctx := c.Request().Context()

	noId, _ := strconv.Atoi(c.Param("no_id"))

	cookie, err := c.Cookie("jwt")
	if err != nil {
		return helper.WriteResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", nil)
	}

	err = d.usecase.DeleteShoesByNoId(ctx, noId, cookie.Value)
	if err != nil {
		return helper.WriteResponse(c, http.StatusNotFound, err.Error(), nil)

	}

	return helper.WriteResponse(c, http.StatusOK, "successfuly deleted data", nil)
}
