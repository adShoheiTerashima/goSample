package main

import (
	"fmt"

	"github.com/adShoheiTerashima/goSample/app/infra"
	"github.com/adShoheiTerashima/goSample/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// debug用。requestとresponseが出ます
func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}

func main() {
	e := echo.New()

	// .env読み込み
	if err := utils.LoadEnv(); err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// debug用。requestとresponseが出ます
	e.Use(middleware.BodyDump(bodyDumpHandler))

	// 最初に発行しておきたいやつ
	db, err := infra.DBCon()
	if err != nil {
		e.Logger.Fatal(err)
	}
	Route(e, db)
	e.Logger.Fatal(e.Start(":1323"))
}
