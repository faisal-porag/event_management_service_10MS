package state

import (
	"event_management_service_10MS/config"
	"event_management_service_10MS/db_repo"
	"github.com/rs/zerolog/log"
)

type State struct {
	Cfg          *config.Config
	DbRepository *db_repo.MySQLRepository
}

func NewState(cfg *config.Config) *State {
	db, err := db_repo.NewMySQLRepository(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("pg connection error")
	} else {
		log.Info().Msg("MYSQL DB connection success")
	}

	return &State{
		Cfg:          cfg,
		DbRepository: db,
	}
}
