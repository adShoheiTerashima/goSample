package domain

import "time"

type times struct {
	CreatedAt time.Time `json:"craeted_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Token フロントとの認証に利用する構造体
type Token struct {
	APIToken     string    `json:"api_token"`
	TokenExpires time.Time `json:"token_expires"`
}

// Account hotel managerのアカウント管理用の構造体
type Account struct {
	HotelManagerID        int64     `json:"hotel_manager_id"`
	ClientCompanyID       int64     `json:"client_company_id"`
	PropertyID            int64     `json:"property_id"`
	WholesalerID          int64     `json:"wholesaler_id"`
	FirstNameEnc          string    `json:"first_name_enc"`
	LastNameEnc           string    `json:"last_name_enc"`
	EmailEnc              string    `json:"email_enc"`
	UsernameEnc           string    `json:"username_enc"`
	PasswordEnc           string    `json:"password_enc"`
	IsPrimary             bool      `json:"is_primary"`
	MasterEditFlg         bool      `json:"master_edit_flg"`
	SettlementNeedFlg     bool      `json:"settlement_need_flg"`
	ClosingDate           int       `json:"closing_date"`
	PaymentDateMonthLater int       `json:"payment_date_month_later"`
	DelFlg                bool      `json:"del_flg"`
	LoginAt               time.Time `json:"login_at"`
	LoginFailedNum        int       `json:"login_failed_num"`
	RememberToken         string    `json:"remember_token"`
	Token                 Token
	Times                 times
}

// IAccountUsecase アカウント関連のusecaseのインターフェース
type IAccountUsecase interface {
	Login(username string, password string) (string, error)
	Logout(hotelManagetID int64) error
	CheckToken(token string) error
}

// IAccountRepository アカウント関連のusecaseのインターフェース
type IAccountRepository interface {
}

// IAccountAPI アカウント関連のusecaseのインターフェース
type IAccountAPI interface {
}
