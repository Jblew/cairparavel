package messenger

// Source: https://chatbotslife.com/facebook-messenger-bot-a-tutorial-in-go-c2aa13b50110

// InputMessage the message we get from Messenger
type InputMessage struct {
	Object string `json:"object"`
	Entry  []struct {
		ID      string `json:"id"`
		Time    int64  `json:"time"`
		Payload struct {
			Title   string `json:"title"`
			Payload string `json:"payload"`
		} `json:"payload"`
		Messaging []struct {
			Sender struct {
				ID string `json:"id"`
			} `json:"sender"`
			Recipient struct {
				ID string `json:"id"`
			} `json:"recipient"`
			Timestamp int64 `json:"timestamp"`
			Message   struct {
				Mid  string `json:"mid"`
				Text string `json:"text"`
				Nlp  struct {
					Entities struct {
						Sentiment []struct {
							Confidence float64 `json:"confidence"`
							Value      string  `json:"value"`
						} `json:"sentiment"`
						Greetings []struct {
							Confidence float64 `json:"confidence"`
							Value      string  `json:"value"`
						} `json:"greetings"`
					} `json:"entities"`
					DetectedLocales []struct {
						Locale     string  `json:"locale"`
						Confidence float64 `json:"confidence"`
					} `json:"detected_locales"`
				} `json:"nlp"`
			} `json:"message"`
			Referral struct {
				Ref    string `json:"ref"`
				Source string `json:"source"`
				Type   string `json:"type"`
			} `json:"referral"`
		} `json:"messaging"`
	} `json:"entry"`
}

// Recipient the recipient of our message
type Recipient struct {
	ID string `json:"id"`
}

// Message the message to send it its basic
type Message struct {
	Text string `json:"text,omitempty"`
}

// Button the button
type Button struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Payload string `json:"payload,omitempty"`
	URL     string `json:"url,omitempty"`
}

// Element an element
type Element struct {
	Title         string        `json:"title,omitempty"`
	Subtitle      string        `json:"subtitle,omitempty"`
	ImageURL      string        `json:"image_url,omitempty"`
	DefaultAction DefaultAction `json:"default_action,omitempty"`
	Buttons       []Button      `json:"buttons,omitempty"`
}

// DefaultAction is default action
type DefaultAction struct {
	Type                string `json:"type,omitempty"`
	URL                 string `json:"url,omitempty"`
	WebViewHeightRation string `json:"webview_height_ratio,omitempty"`
}

// Attachment the attachment to send (custom)
type Attachment struct {
	Attachment struct {
		Type    string `json:"type,omitempty"`
		Payload struct {
			TemplateType string    `json:"template_type,omitempty"`
			Elements     []Element `json:"elements,omitempty"`
		} `json:"payload,omitempty"`
	} `json:"attachment,omitempty"`
}

// ResponseAttachment full response
type ResponseAttachment struct {
	Recipient Recipient  `json:"recipient"`
	Message   Attachment `json:"message,omitempty"`
}

// ResponseMessage full response
type ResponseMessage struct {
	Recipient Recipient `json:"recipient"`
	Message   Message   `json:"message,omitempty"`
}
