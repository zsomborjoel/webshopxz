package session

import (
	"os"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/rs/zerolog/log"
)

func SetupStore() cookie.Store {
	storekey := os.Getenv("SESSION_STORE_KEY")
	if storekey == "" {
		log.Error().Msg("SESSION_STORE_KEY environment variable is not set")
	}

	return cookie.NewStore([]byte(storekey))
}
