package main

import (
	config "github.com/michaelCHU95/auth-hub/initializer"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var SqlDB *gorm.DB

func init() {
	config.LoadEnvVariables()
	SqlDB = config.ConnectToDB()
	config.SyncDatabase(SqlDB)
}

func main() {
	app := gin.Default()

	app.Run()
}
