package email

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func getConfig() (*oauth2.Config, error) {
	secretPath := os.Getenv("SECRET_PATH")

	b, err := os.ReadFile(secretPath)
	if err != nil {
		return nil, fmt.Errorf("Unable to read client secret file: %w", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailComposeScope)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse client secret file to config: %w", err)
	}

	return config, nil
}

// Retrieve a token, saves the token, then returns the generated client.
// The file token.json stores the user's access and refresh tokens, and is
// created automatically when the authorization flow completes for the first time.
func getClient(config *oauth2.Config) *http.Client {
	log.Debug().Msg(fmt.Sprintf("email.getClient started"))

	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		log.Info().Msg("Getting token from web")

		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}

	cli := config.Client(context.Background(), tok)

	log.Debug().Msg(fmt.Sprintf("email.getClient done"))

	return cli
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	log.Debug().Msg(fmt.Sprintf("email.getTokenFromWeb started"))

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	log.Info().Msg(fmt.Sprintf("Go to the following link in your browser then type the "+
		"authorization code: %s", authURL))

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Error().Err(err).Msg("Unable to read authorization code")
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Error().Err(err).Msg("Unable to retrieve token from web")
	}

	log.Debug().Msg(fmt.Sprintf("email.getTokenFromWeb done"))

	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	log.Debug().Msg(fmt.Sprintf("email.tokenFromFile started"))

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)

	log.Debug().Msg(fmt.Sprintf("email.tokenFromFile started"))

	return tok, nil
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	log.Debug().Msg(fmt.Sprintf("email.saveToken started"))

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Error().Err(err).Msg("Unable to cache oauth token")
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)

	log.Debug().Msg(fmt.Sprintf("email.saveToken done"))
}
