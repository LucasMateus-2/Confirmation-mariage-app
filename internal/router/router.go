package router

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lucas/confirmation-mariage-app/internal/handler"
	"github.com/lucas/confirmation-mariage-app/internal/middleware"
)

func Setup(
	authHandler *handler.AuthHandler,
	guestHandler *handler.GuestHandler,
) *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins(),
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Público
	router.POST("/auth/login", authHandler.Login)
	router.GET("/confirm/:id", guestHandler.GetForConfirmation)
	router.PATCH("/guests/:id/confirm", guestHandler.Confirm)
	router.GET("/guests/search", guestHandler.SearchByName)
	// Protegido — apenas os noivos
	protected := router.Group("/")
	protected.Use(middleware.Auth())
	{
		protected.GET("/guests", guestHandler.ListAll)
		protected.GET("/guests/:id", guestHandler.GetByID)
		protected.GET("/dashboard", guestHandler.Dashboard)
	}

	return router
}

func allowedOrigins() []string {
	origins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if origins == "" {
		return []string{"http://localhost:5173"}
	}
	return strings.Split(origins, ",")
}
