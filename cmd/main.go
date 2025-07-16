package main

import (
	"test-naga-exchange/config"
	"test-naga-exchange/controller"
	"test-naga-exchange/middleware"
	"test-naga-exchange/model"
	"test-naga-exchange/repository"
	"test-naga-exchange/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}
	db := config.InitDB()

	db.AutoMigrate(
		&model.User{},
		&model.Transaction{},
	)

	r := gin.Default()

	userRepo := repository.NewUserRepository(config.DB)
	txRepo := repository.NewTransactionRepository(config.DB)
	authService := service.NewAuthService(userRepo)
	txService := service.NewTransactionService(txRepo)
	authHandler := controller.NewAuthHandler(authService)
	txHandler := controller.NewTransactionHandler(txService)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(db))
	{
		auth.GET("/transaction", txHandler.Get)
		auth.POST("/transaction/process", txHandler.Process)
	}

	r.Run(":8081")
}
