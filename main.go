package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/WilkerAlves/genealogy/infra/http"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal().Msg("Environment variable not defined")
	}
	envProps := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
		"DB_USER",
		"DB_PASS",
		"SERVER_PORT",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			log.Fatal().Msg(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339Nano,
	})

	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	http.StartServer(rootCtx)
}
