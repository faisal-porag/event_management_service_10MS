package main

import (
	"event_management_service_10MS/config"
	"event_management_service_10MS/httpserver"
	"event_management_service_10MS/state"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {

	envLoadErr := godotenv.Load()
	if envLoadErr != nil {
		log.Fatal().Err(envLoadErr).Msg("env load error")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Config parsing failed")
	}

	appState := state.NewState(cfg)
	httpserver.Serve(appState)
}
