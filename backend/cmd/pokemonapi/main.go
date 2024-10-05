package main

import (
	"internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middleware.AuthMiddleware())
}
