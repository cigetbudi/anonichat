package api

import (
	"anonichat/models"
	"anonichat/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMessages(ctx *gin.Context) {
	mes, err := models.GetAllMessages()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, mes)
}

func GetAllMessagesByUser(ctx *gin.Context) {
	user_id, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mes, err := models.GetMessagesByUserID(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, mes)
}

func CreateMessage(ctx *gin.Context) {
	m := models.Message{}
	if err := ctx.ShouldBind(&m); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "mohon untuk melengkapi semua isian"})
		return
	}
	err := models.CreateMessage(&m)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "berhasil membuat postingan baru"})
}

func DeleteMessage(ctx *gin.Context) {
	user_id, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.DeleteMessage(uint(id), user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "berhasil menghapus postingan"})
}

func LikePost(ctx *gin.Context) {
	_, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.AddLike(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "berhasil menyukai postingan"})

}

func UnlikePost(ctx *gin.Context) {
	_, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.UnLike(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "berhasil menghapus like postingan"})

}
