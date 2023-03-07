package main

import (
	"github.com/michaelCHU95/auth-hub/api"
	"github.com/michaelCHU95/auth-hub/initializer"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var SqlDB *gorm.DB

func init() {
	initializer.LoadEnvVariables()
	SqlDB = initializer.ConnectToDB()
	initializer.SyncDatabase(SqlDB)
}

func main() {
	app := gin.Default()

	v1 := app.Group("/api")
	api.NewRouter(v1, SqlDB)

	app.Run()
}
