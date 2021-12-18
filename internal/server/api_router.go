package server

import (
	"github.com/gin-gonic/gin"
	"mvc/internal/app/api/user"
)

func ApiRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", user.CreateUser)
		userGroup.POST("/login", user.Login)
		userGroup.POST("/logout", user.Logout)
	}
}
