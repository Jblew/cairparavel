// Based on https://chatbotslife.com/facebook-messenger-bot-a-tutorial-in-go-c2aa13b50110

package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/Jblew/cairparavel/functions/app/apps/messengerapp"
	"github.com/Jblew/cairparavel/functions/app/lib/messenger"
	"github.com/Jblew/cairparavel/functions/util"
)

// FnMessengerWebhook is FB messenger webhook
func FnMessengerWebhook(resp http.ResponseWriter, request *http.Request) {
	logWithCode := func(code int) func(format string, v ...interface{}) {
		return func(format string, v ...interface{}) {
			str := fmt.Sprintf(format, v...)
			log.Printf(str)
			resp.WriteHeader(code)
			resp.Write([]byte(str))
		}
	}

	opts := util.FunctionHandlerOpts{
		Name:       "FnMessengerWebhook",
		LogErrorFn: logWithCode(400),
		LogPanicFn: logWithCode(200),
		LogDoneFn:  logWithCode(200),
	}
	util.FunctionHandler(opts, func() error {
		var messengerInstance *messenger.Messenger
		container.Make(&messengerInstance)

		if request.Method == "GET" {
			u, _ := url.Parse(request.RequestURI)
			values, _ := url.ParseQuery(u.RawQuery)
			token := values.Get("hub.verify_token")
			if messengerInstance.IsVerifyTokenCorrect(token) {
				resp.WriteHeader(200)
				resp.Write([]byte(values.Get("hub.challenge")))
				return nil
			}
			resp.WriteHeader(400)
			resp.Write([]byte(`Bad token`))
			return nil
		}

		// Anything that reaches here is POST.
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return err
		}

		// Parse message into the Message struct
		var message messenger.InputMessage
		err = json.Unmarshal(body, &message)
		if err != nil {
			return err
		}
		return messengerapp.OnMessengerInputMessage(message, container)
	})
}
