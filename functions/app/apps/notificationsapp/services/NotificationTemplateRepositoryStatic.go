package services

import (
	"fmt"
)

var templates = map[string]string{
	"comment_added":                "New comment on event {{ .event.name }}",
	"event_voting_started":         "Voting started on event {{ .event.name }}",
	"event_members_signup_started": "Signup started on event {{ .event.name }}",
	"event_members_signup_closed":  "Signup closed on event {{ .event.name }}",
	"event_started":                "Event {{ .event.name }} started",
	"event_cancelled":              "Event {{ .event.name }} cancelled",
	"messenger_respond":            "This is a bot account. Please visit https://cairparavelapp.web.app to check information on your events",
	"messenger_welcome":            "Hi {{ .user.DisplayName }}! You have just enabled updates on our events.",
	"member_signed_in":             "New person signed in on event {{ .event.name }}",
	"member_signed_out":            "One person signed out on event {{ .event.name }}",
	"event_voted":                  "New vote on event {{ .event.name }}",
	"event_vote_deleted":           "Vote deleted on event {{ .event.name }}",
}

// NotificationTemplateRepositoryStatic implements NotificationTemplateRepository
type NotificationTemplateRepositoryStatic struct {
}

func (repo *NotificationTemplateRepositoryStatic) GetTemplate(Name string) (NotificationTemplate, error) {
	if val, ok := templates[Name]; ok {
		return NotificationTemplate{
			Name:     Name,
			Template: val,
		}, nil
	}
	return NotificationTemplate{}, fmt.Errorf("Template with name %s doesnt exist", Name)
}

// NotificationTemplate a template for notification
type NotificationTemplate struct {
	Name     string
	Template string
}
