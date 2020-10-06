package functions

import (
	"net/http"
	"net/url"
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
}
