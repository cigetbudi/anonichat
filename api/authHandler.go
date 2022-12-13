package api

import (
	"anonichat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	vm := models.Register{}
	if err := ctx.ShouldBind(&vm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "mohon untuk melengkapi semua isian"})
		return
	}

	u := models.User{}
	u.Username = vm.Username
	u.Email = vm.Email
	u.Name = vm.Name
	u.Password = vm.Password

	_, err := u.SaveUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "berhasil melakukan registrasi"})

}

func Login(ctx *gin.Context) {
	vm := models.Login{}
	if err := ctx.ShouldBind(&vm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Username = vm.Username
	u.Password = vm.Password

	token, err := models.LoginCheck(u.Username, u.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username atau password tidak sesuai"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})

}
func Logout(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "API Logut",
	})

}
