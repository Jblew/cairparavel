// Based on https://chatbotslife.com/facebook-messenger-bot-a-tutorial-in-go-c2aa13b50110

package functions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/Jblew/cairparavel/functions/app/messenger"
)

// FnMessengerWebhook is FB messenger webhook
func FnMessengerWebhook(resp http.ResponseWriter, request *http.Request) {
	verifyToken := application.Config.Messenger.VerifyToken

	if request.Method == "GET" {
		u, _ := url.Parse(request.RequestURI)
		values, _ := url.ParseQuery(u.RawQuery)
		token := values.Get("hub.verify_token")
		if token == verifyToken {
			resp.WriteHeader(200)
			resp.Write([]byte(values.Get("hub.challenge")))
			return
		}
		resp.WriteHeader(400)
		resp.Write([]byte(`Bad token`))
		return
	}

	// Anything that reaches here is POST.
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Failed parsing body: %s", err)
		resp.WriteHeader(400)
		resp.Write([]byte("An error occurred"))
		return
	}

	// Parse message into the Message struct
	var message messenger.InputMessage
	err = json.Unmarshal(body, &message)
	if err != nil {
		log.Printf("Failed unmarshalling message: %s", err)
		resp.WriteHeader(400)
		resp.Write([]byte("An error occurred"))
		return
	}
	// Find messages
	log.Printf("Message: %#v", message)
	for _, entry := range message.Entry {
		if len(entry.Messaging) == 0 {
			log.Printf("No messages")
			resp.WriteHeader(400)
			resp.Write([]byte("An error occurred"))
			return
		}
		for _, event := range entry.Messaging {
			log.Printf("Event: %#v", event)
			if len(event.Referral.Ref) > 0 {
				log.Printf("REFERRAL FOUND: %s", event.Referral.Ref)
			}
			err = handleMessage(event.Sender.ID, event.Message.Text)
			if err != nil {
				log.Printf("Failed sending message: %s", err)
				resp.WriteHeader(400)
				resp.Write([]byte("An error occurred"))
				return
			}
		}
	}
}

// Handles messages
func handleMessage(senderID, message string) error {
	accessToken := application.Config.Messenger.AccessToken

	if len(message) == 0 {
		return errors.New("no message found")
	}
	response := messenger.ResponseMessage{
		Recipient: messenger.Recipient{
			ID: senderID,
		},
		Message: messenger.Message{
			Text: "Hello",
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
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed doing request: %s", err)
		return err
	}
	log.Printf("MESSAGE SENT?\n%#v", res)
	return nil
}
