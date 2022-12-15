package api

import (
	"anonichat/middlwares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to AnoniCHAT OPENAPI",
		})
	})

	auth := r.Group("/auth")
	auth.POST("/register", Register)
	auth.POST("/login", Login)
	auth.POST("/logout", Logout)

	protected := r.Group("/admin")
	protected.Use(middlwares.JwtAuthMiddleware())
	protected.GET("/user", CurrentLogin)

	urArea := r.Group("/urArea")
	urArea.Use(middlwares.JwtAuthMiddleware())
	urArea.GET("/GetAllPosts", GetAllMessages)
	urArea.GET("/GetAllPostsByUser", GetAllMessagesByUser)
	urArea.POST("/CreatePost", CreateMessage)
	urArea.POST("/LikePost/:id", LikePost)
	urArea.POST("/UnlikePost/:id", UnlikePost)
	urArea.DELETE("/DeletePost/:id", DeleteMessage)

	return r
}
