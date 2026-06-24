package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/lucas/confirmation-mariage-app/internal/handler"
	"github.com/lucas/confirmation-mariage-app/internal/infra"
	"github.com/lucas/confirmation-mariage-app/internal/repository"
	"github.com/lucas/confirmation-mariage-app/internal/router"
	"github.com/lucas/confirmation-mariage-app/internal/service"
	"github.com/lucas/confirmation-mariage-app/pkg/logger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Info("aviso: .env não encontrado, usando variáveis de ambiente do sistema")
	}

	logger.Init()
	logger.Info("iniciando servidor...")

	db, err := infra.NewDB()
	if err != nil {
		logger.Fatal(err, "erro ao conectar no banco")
	}
	defer db.Close()

	logger.Info("banco de dados conectado")

	// Repositories
	userRepo := repository.NewUserRepository(db)
	guestRepo := repository.NewGuestRepository(db)

	// Services
	authService := service.NewAuthService(userRepo)
	guestService := service.NewGuestService(guestRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)
	guestHandler := handler.NewGuestHandler(guestService)

	// Router
	r := router.Setup(authHandler, guestHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.With().Str("port", port).Msg("servidor pronto")

	if err := r.Run(":" + port); err != nil {
		logger.Fatal(err, "erro ao iniciar servidor")
	}
}
