// Based on https://chatbotslife.com/facebook-messenger-bot-a-tutorial-in-go-c2aa13b50110

package functions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/Jblew/cairparavel/functions/app/apps/messengerapp"
	"github.com/Jblew/cairparavel/functions/app/lib/messenger"
)

// FnMessengerWebhook is FB messenger webhook
func FnMessengerWebhook(resp http.ResponseWriter, request *http.Request) {
	var messengerInstance *messenger.Messenger
	container.Make(&messengerInstance)

	if request.Method == "GET" {
		u, _ := url.Parse(request.RequestURI)
		values, _ := url.ParseQuery(u.RawQuery)
		token := values.Get("hub.verify_token")
		if messengerInstance.IsVerifyTokenCorrect(token) {
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
	err = messengerapp.OnMessengerInputMessage(message, container)
	if err != nil {
		log.Printf("Failed unmarshalling message: %s", err)
		resp.WriteHeader(400)
		resp.Write([]byte("An error occurred"))
		return
	}
}
