package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	zerolog.TimeFieldFormat = time.RFC3339

	env := os.Getenv("APP_ENV")
	if env == "production" {
		// JSON puro em produção (melhor pra coleta de logs)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	} else {
		// Pretty print colorido em desenvolvimento
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "15:04:05",
		}).With().Timestamp().Caller().Logger()
	}
}
func Info(msg string) {
	log.Info().Msg(msg)
}

func Error(err error, msg string) {
	log.Error().Err(err).Msg(msg)
}

func Fatal(err error, msg string) {
	log.Fatal().Err(err).Msg(msg)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

// With retorna o logger global pra uso com campos customizados
// Ex: logger.With().Str("user_id", "123").Msg("login realizado")
func With() *zerolog.Event {
	return log.Info()
}
