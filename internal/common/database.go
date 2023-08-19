package common

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var DB *sqlx.DB

func Init() *sqlx.DB {
	conn := os.Getenv("DB_CONNECTION")

	var err error
	DB, err = sqlx.Connect("postgres", conn)
	if err != nil {
		log.Error().Stack().Err(err).Msg("db err: (Init)")
	}

	return DB
}

func GetDB() *sqlx.DB {
	return DB
}
