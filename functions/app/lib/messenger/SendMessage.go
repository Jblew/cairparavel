// Based on https://chatbotslife.com/facebook-messenger-bot-a-tutorial-in-go-c2aa13b50110

package messenger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// SendMessage sends FB messenger message
func (messenger *Messenger) SendMessage(recipient Recipient, text string) error {
	accessToken := messenger.Config.AccessToken

	response := ResponseMessage{
		Recipient: recipient,
		Message: Message{
			Text: text,
		},
	}
	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("Marshal error: %s", err)
		return err
	}
	uri := "https://graph.facebook.com/v2.6/me/messages"
	uri = fmt.Sprintf("%s?access_token=%s", uri, accessToken)
	log.Printf("URI: %s", uri)
	req, err := http.NewRequest(
		"POST",
		uri,
		bytes.NewBuffer(data),
	)
	if err != nil {
		log.Printf("Failed making request: %s", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
