package email

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/zsomborjoel/workoutxz/internal/common"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

const EmailTitle = "Workoutxz account confirmation"
const User = "me"
const EndpointDummy = "$?"

func SendEmail(to string, token string) error {
	log.Debug().Msg("email.SendEmail called")

	from := os.Getenv("GMAIL_ACCOUNT")

	config, err := getConfig()
	if err != nil {
		return fmt.Errorf("Unable to parse client secret file to config: %w", err)
	}
	client := getClient(config)

	srv, err := gmail.NewService(
		context.Background(),
		option.WithScopes(gmail.GmailSendScope),
		option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("Unable to retrieve Gmail client: %w", err)
	}
	log.Debug().Msg("Service created")

	html, err := getEmailHtml(token)
	if err != nil {
		return err
	}

	msgStr := "From:" + from + "\r\n"
	msgStr += "To: " + to + "\r\n"
	msgStr += "Subject: " + EmailTitle + "\r\n"
	msgStr += "Content-Type: text/html; charset=\"UTF-8\"\r\n"
	msgStr += html

	msg := []byte(msgStr)
	gMessage := &gmail.Message{Raw: base64.URLEncoding.EncodeToString(msg)}

	_, err = srv.Users.Messages.Send(User, gMessage).Do()
	if err != nil {
		return fmt.Errorf("Email send failed: %w", err)
	}
	log.Info().Msg(fmt.Sprintf("Email sent to: %s", to))

	return nil
}

func getEmailHtml(token string) (string, error) {
	htmlPath := os.Getenv("EMAIL_HTML_PATH")
	url := os.Getenv("APP_URL")
	endpoint := url + "/api/auth" + common.ConfirmRegistrationEndpoint + "?token=" + token

	html, err := os.ReadFile(htmlPath)
	if err != nil {
		return "", fmt.Errorf("Html Path file read failed: %w", err)
	}

	replacedHtml := strings.ReplaceAll(string(html), EndpointDummy, endpoint)

	return replacedHtml, nil
}
