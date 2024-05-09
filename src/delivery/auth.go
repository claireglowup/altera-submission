package delivery

import (
	"net/http"
	"restapi-altera/src/helper"
	"restapi-altera/src/model/request"
	"time"

	"github.com/labstack/echo/v4"
)

func (d *delivery) RegisterHandler(c echo.Context) error {

	var user request.UserRegister

	ctx := c.Request().Context()

	if err := c.Bind(&user); err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err := c.Validate(user)
	if err != nil {
		return helper.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	err = d.usecase.Register(ctx, user)
	if err != nil {
		return helper.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	return helper.WriteResponse(c, http.StatusOK, "succes register user", nil)

}

func (d *delivery) Login(c echo.Context) error {

	var user request.UserLogin

	ctx := c.Request().Context()

	if err := c.Bind(&user); err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	err := c.Validate(user)
	if err != nil {
		return helper.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	token, err := d.usecase.Login(ctx, user)
	if err != nil {
		return helper.WriteResponse(c, http.StatusNotFound, err.Error(), nil)
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		Path:     "/user",
		HttpOnly: true,
	}

	c.SetCookie(cookie)

	return helper.WriteResponse(c, http.StatusOK, "succesfully login", nil)

}

func (d *delivery) Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		Path:     "/user",
		HttpOnly: true,
	}

	c.SetCookie(cookie)

	return helper.WriteResponse(c, http.StatusOK, "you are logout", nil)

}
