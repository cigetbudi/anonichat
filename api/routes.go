package api

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	auth.POST("/register", Register)
	auth.POST("/login", Login)
	auth.POST("/logout", Logout)

	return r
}
