package infra

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBCon DB読み込み処理
func DBCon() (*gorm.DB, error) {

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	protocol := fmt.Sprintf("tcp(%s:%s)", host, port)

	connectInfo := fmt.Sprintf("%s:%s@%s/%s", username, password, protocol, database)
	DB, err := gorm.Open(mysql.Open(connectInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return DB, nil
}
