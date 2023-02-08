package main

import (
	"be13/ca/config"
	"be13/ca/factory"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectToDB()
	defer db.Close()

	e := gin.New()
	e.Use(gin.Recovery())

	factory.InitFactory(e, db)
	e.Run(fmt.Sprintf(":%d", 8080))

}
