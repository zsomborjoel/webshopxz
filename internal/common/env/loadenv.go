package env

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Stack().Msg("Error loading .env file")
	}
}
