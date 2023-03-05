package api

import (
	"github.com/gin-gonic/gin"
)

func HandleSignUp() gin.HandlerFunc {
	fn := func(c *gin.Context) {}

	return gin.HandlerFunc(fn)
}

func HandleLogin() {}

func HandlePasswordReset() {}

func HandleReprovision() {}
