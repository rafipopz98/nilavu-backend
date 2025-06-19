package auth_routes

import (
	"github.com/gin-gonic/gin"
	auth_controller "github.com/nilavu-backend/src/api/controller/auth"
)

func AuthRoutes (router *gin.RouterGroup){
	// signup
	router.POST("/signup", auth_controller.SignUp);
	// login
	router.POST("/login", auth_controller.Login);

}
