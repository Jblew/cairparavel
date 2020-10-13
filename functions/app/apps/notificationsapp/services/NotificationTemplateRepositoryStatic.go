package services

import (
	"fmt"

	"github.com/Jblew/cairparavel/functions/app/apps/notificationsapp/notificationsdomain"
)

var templates = map[string]string{
	"comment_added":                "New comment on event {{ .event.name }}",
	"event_voting_started":         "Voting started on event {{ .event.name }}",
	"event_members_signup_started": "Signup started on event {{ .event.name }}",
	"event_members_signup_closed":  "Signup closed on event {{ .event.name }}",
	"event_started":                "Event {{ .event.name }} started",
	"event_cancelled":              "Event {{ .event.name }} cancelled",
	"messenger_respond":            "This is a bot account. Please visit https://cairparavelapp.web.app to check information on your events",
	"messenger_welcome":            "Hi! You have just enabled updates on our events. Visit https://cairparavelapp.web.app to subscribe or unsubscribe particular events.",
	"member_signed_in":             "New person signed in on event {{ .event.name }}",
	"member_signed_out":            "One person signed out on event {{ .event.name }}",
	"event_voted":                  "New vote on event {{ .event.name }}",
	"event_vote_deleted":           "Vote deleted on event {{ .event.name }}",
	"event_created":                "New event created {{ .event.name }}",
	"event_deleted":                "Event {{ .event.name }} was deleted",
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
