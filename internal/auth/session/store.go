package session

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/rs/zerolog/log"
)

func SetupStore() cookie.Store {
	storekey := os.Getenv("SESSION_STORE_KEY")
	if storekey == "" {
		log.Error().Msg("SESSION_STORE_KEY environment variable is not set")
	}

	maxAgeValue := os.Getenv("SESSION_STORE_MAX_AGE")
	if maxAgeValue == "" {
		maxAgeValue = "86400" // 1 day in seconds
		log.Error().Msg("SESSION_STORE_MAX_AGE environment variable is not set")
	}

	maxAge, err := strconv.Atoi(maxAgeValue)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("SESSION_STORE_MAX_AGE can not be converted to int: %s", maxAgeValue))
	}

	secureValue := os.Getenv("SESSION_STORE_SECURE")
	if secureValue == "" {
		secureValue = "false" // Set to true if using HTTPS
		log.Error().Msg("SESSION_STORE_SECURE environment variable is not set")
	}

	secure, err := strconv.ParseBool(secureValue)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("SESSION_STORE_SECURE can not be converted to bool: %s", secureValue))
	}

	domain := os.Getenv("SESSION_STORE_DOMAIN")
	if domain == "" {
		log.Error().Msg("SESSION_STORE_DOMAIN environment variable is not set")
	}

	store := cookie.NewStore([]byte(storekey))
	store.Options(sessions.Options{
		MaxAge:   maxAge,
		HttpOnly: true, // Against XSS by fetching document.cookie with JS
		Secure:   secure,
		Domain:   domain,
	})

	return store
}
