package user

import (
	"net/http"

	"github.com/ebilsanta/social-network/backend/complex_services/user_service/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var user models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
