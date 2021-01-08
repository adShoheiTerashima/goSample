package infra

import (
	"github.com/adShoheiTerashima/goSample/domain"
)

type accountAPI struct {
}

// NewAccountAPI hoge
func NewAccountAPI() domain.IAccountAPI {
	return &accountAPI{}
}
