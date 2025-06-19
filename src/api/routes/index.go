package routes

import (
	"github.com/gin-gonic/gin"
	auth_routes "github.com/nilavu-backend/src/api/routes/auth"
)
func InitRouters () *gin.Engine{
	router := gin.Default();
	apiGroup := router.Group("/api")

	authGroup := apiGroup.Group("/auth")
	
	auth_routes.AuthRoutes(authGroup)


	return router;
}