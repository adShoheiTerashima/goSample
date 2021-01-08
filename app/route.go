package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	account "github.com/adShoheiTerashima/goSample/account/handler"
)

// Route hoge
func Route(e *echo.Echo, db *gorm.DB) {
	accountHandler := account.NewAccountHandler(db)
	e.POST("/login", accountHandler.Login)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hoge")
	})
}
