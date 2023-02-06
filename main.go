package main

import (
	"be13/ca/config"
	"be13/ca/factory"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.ConnectToDB()

	e := echo.New()

	factory.InitFactory(e, db)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 8080)))

}
