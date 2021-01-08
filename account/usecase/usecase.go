package usecase

import (
	"github.com/adShoheiTerashima/goSample/account/infra"
	"github.com/adShoheiTerashima/goSample/domain"
	"gorm.io/gorm"
)

type accountUsecase struct {
	ARepository domain.IAccountRepository
	AAPI        domain.IAccountAPI
}

// NewAccountUsecase hoge
func NewAccountUsecase(db *gorm.DB) domain.IAccountUsecase {
	return &accountUsecase{
		ARepository: infra.NewAccountRepository(db),
		AAPI:        infra.NewAccountAPI(),
	}
}

/*

	Login(username string, password string) (string, error)
	Logout(hotelManagetID int64) error
	CheckToken(token string) error
*/
func (a *accountUsecase) Login(username string, password string) (string, error) {
	return "", nil
}

func (a *accountUsecase) Logout(hotelManagetID int64) error {
	return nil
}

func (a *accountUsecase) CheckToken(token string) error {
	return nil
}
