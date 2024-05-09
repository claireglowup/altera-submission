package delivery

import "github.com/labstack/echo/v4"

func (d *delivery) Auth(e *echo.Group) {
	e.POST("/register", d.RegisterHandler)
	e.POST("/login", d.Login)
	e.POST("/logout", d.Logout)
}

func (d *delivery) User(e *echo.Group) {
	e.GET("", d.GetUserByID)
	e.GET("/shoes", d.GetUserAllShoes)
	e.GET("/shoes/:no_id", d.GetUserShoesByNoId)
	e.POST("/shoes", d.AddNewShoes)
	e.PUT("/shoes/:no_id", d.UpdateShoesByNoId)
	e.DELETE("/shoes/:no_id", d.DeleteShoesByNoId)
}
