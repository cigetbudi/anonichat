package api

import (
	"anonichat/middlwares"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	auth.POST("/register", Register)
	auth.POST("/login", Login)
	auth.POST("/logout", Logout)

	protected := r.Group("/admin")
	protected.Use(middlwares.JwtAuthMiddleware())
	protected.GET("/user", CurrentLogin)

	return r
}
