package infra

import (
	"github.com/adShoheiTerashima/goSample/domain"
	"gorm.io/gorm"
)

type accountRepository struct {
}

// NewAccountRepository hoge
func NewAccountRepository(db *gorm.DB) domain.IAccountRepository {
	return &accountRepository{}
}
