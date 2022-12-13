package main

import (
	"anonichat/database"

	"github.com/gin-gonic/gin"
)

func init() {
	database.InitDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "API JALAN",
		})
	})

	r.Run(":4545")

}
