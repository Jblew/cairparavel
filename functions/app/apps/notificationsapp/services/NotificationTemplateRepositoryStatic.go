package services

import (
	"fmt"

	"github.com/Jblew/cairparavel/functions/app/apps/notificationsapp/notificationsdomain"
)

var templates = map[string]string{
	"comment_added":                "New comment on event {{ .event.Name }}",
	"event_voting_started":         "Voting started on event {{ .event.Name }}",
	"event_members_signup_started": "Signup started on event {{ .event.Name }}",
	"event_members_signup_closed":  "Signup closed on event {{ .event.Name }}",
	"event_started":                "Event {{ .event.Name }} started",
	"event_cancelled":              "Event {{ .event.Name }} cancelled",
	"messenger_respond":            "This is a bot account. Please visit https://cairparavelapp.web.app to check information on your events",
	"messenger_welcome":            "Hi {{ .user.DisplayName }}! You have just enabled updates on our events.",
	"member_signed_in":             "New person signed in on event {{ .event.Name }}",
	"member_signed_out":            "One person signed out on event {{ .event.Name }}",
	"event_voted":                  "New vote on event {{ .event.Name }}",
	"event_vote_deleted":           "Vote deleted on event {{ .event.Name }}",
	"event_created":                "New event created {{ .event.Name }} {{ event.Name }}",
}

// NotificationTemplateRepositoryStatic implements NotificationTemplateRepository
type NotificationTemplateRepositoryStatic struct {
}

// GetTemplate returns staticly stored template
func (repo *NotificationTemplateRepositoryStatic) GetTemplate(Name string) (notificationsdomain.NotificationTemplate, error) {
	if val, ok := templates[Name]; ok {
		return notificationsdomain.NotificationTemplate{
			Name:     Name,
			Template: val,
		}, nil
	}
	return notificationsdomain.NotificationTemplate{}, fmt.Errorf("Template with name %s doesnt exist", Name)
}
