package handler

import (
	"net/http"

	"github.com/adShoheiTerashima/goSample/account/usecase"
	"github.com/adShoheiTerashima/goSample/domain"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// AccountHandler  represent the httphandler for account
type AccountHandler struct {
	AUsecase domain.IAccountUsecase
}

// NewAccountHandler will initialize the account resources endpoint
func NewAccountHandler(db *gorm.DB) *AccountHandler {
	return &AccountHandler{
		AUsecase: usecase.NewAccountUsecase(db),
	}
}

// Login hoge
func (a *AccountHandler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
