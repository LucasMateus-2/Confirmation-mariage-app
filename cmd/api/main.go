package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/LucasMateus-2/Confirmation-mariage-app/internal/handler"
	"github.com/LucasMateus-2/Confirmation-mariage-app/internal/middleware"
	"github.com/LucasMateus-2/Confirmation-mariage-app/internal/repository"
	"github.com/LucasMateus-2/Confirmation-mariage-app/internal/service"

	_ "github.com/jackc/pgx/v5/stdlib" // Driver do Postgres
)

func main() {
	jwtSecret := "sua_chave_secreta_super_segura"
	port := ":8080"

	// String de conexão (em produção, use os.Getenv)
	dsn := "postgres://postgres:secret@localhost:5432/login_db?sslmode=disable"

	// 1. Conecta ao Banco de Dados
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir banco de dados: %v", err)
	}
	defer db.Close()

	// Verifica se a conexão está ativa
	if err := db.Ping(); err != nil {
		log.Fatalf("Não foi possível conectar ao banco: %v", err)
	}

	// 2. Injeta o DB no novo repositório Postgres
	userRepo := repository.NewPostgresUserRepository(db)
	authService := service.NewAuthService(userRepo, jwtSecret)
	authHandler := handler.NewAuthHandler(authService)

	// Rotas
	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)

	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(middleware.UserIDKey)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Bem-vindo à área protegida!",
			"user_id": userID,
		})
	})
	http.Handle("/dashboard", middleware.JWTMiddleware(jwtSecret)(protectedMux))

	fmt.Printf("Servidor rodando na porta %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
