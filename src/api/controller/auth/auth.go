package auth_controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	request_auth "github.com/nilavu-backend/src/api/requests/auth"
	"github.com/nilavu-backend/src/api/validation"
)

func SignUp(ctx *gin.Context){
	var body request_auth.SignUP_Request

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	fmt.Println(body)
	
	// validation
	err := validation.ValidateSignUpRequest(body);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"message": "Validation Error",
		})
		return
	}
	// then go services layer
	
	// if all good then return 200
}

func Login(ctx *gin.Context){
	
}