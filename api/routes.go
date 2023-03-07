package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(r *gin.RouterGroup, db *gorm.DB) {
	// Run all routes here

	r.POST("/signup", HandleSignup(db))
	r.POST("/login", HandleLogin(db))
}
