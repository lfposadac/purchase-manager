package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lfposadac/purchase-manager/internal/infraestructure/handler"
	"github.com/lfposadac/purchase-manager/internal/infraestructure/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.POST("/login", handler.LoginHandler)
	authorized := r.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())
	authorized.GET("/me", handler.MeHandler)

	r.Run(":" + os.Getenv("PORT"))

}
