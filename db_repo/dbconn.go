package db_repo

import (
	"event_management_service_10MS/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type MySQLRepository struct {
	db *sqlx.DB
}

func NewMySQLRepository(cfg *config.Config) (*MySQLRepository, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.MySqlUsername, cfg.MySqlDbPassword, cfg.MySqlHost, cfg.MySqlDbPort, cfg.MySqlDatabaseName)

	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Error().Err(err).Msg("MySQL database is down! Please check")
		return nil, err
	}

	return &MySQLRepository{
		db: db,
	}, nil
}
