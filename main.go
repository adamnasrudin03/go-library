package main

import (
	"go-library/config"
	"go-library/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)
func main() {
	defer config.CloseDatabaseConnection(db)
	router := gin.Default()

	authRoutes := router.Group("/api/v1/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}


	router.Run() 
}