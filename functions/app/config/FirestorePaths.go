package config

// FirestorePaths is a list of paths or path returners for firestore
type FirestorePaths struct {
}

func (paths *FirestorePaths) UsersCol() string {
	return "envs/prod/users"
}

func (paths *FirestorePaths) UserDoc(userID string) string {
	return "envs/prod/users/" + userID
}

func (paths *FirestorePaths) UserDoc(userID string) string {
	return "envs/prod/users/" + userID
}

func (paths *FirestorePaths) MessengerNotificationsForUserCol(uid string) string {
	return "envs/prod/notifications/" + uid + "/messenger_queue"
}

func (paths *FirestorePaths) NotificationsHistoryForUserCol(uid string) string {
	return "envs/prod/notifications/" + uid + "/history"
}

func (paths *FirestorePaths) CommentsForEventCol(eventID string) string {
	return "envs/prod/event_comments/" + eventID + "/messages"
}

func (paths *FirestorePaths) EventsCol() string {
	return "envs/prod/events"
}

func (paths *FirestorePaths) SignupsForEventCol(eventID string) string {
	return "envs/prod/events/" + eventID + "/signedMembers"
}

func (paths *FirestorePaths) SignupsForEventForUserDoc(eventID string, userID string) string {
	return "envs/prod/events/" + eventID + "/signedMembers/" + userID
}

func (paths *FirestorePaths) VotesForEventCol(eventID string) string {
	return "envs/prod/events/" + eventID + "/votes"
}

func (paths *FirestorePaths) VotesForEventForUserDoc(eventID string, userID) string {
	return "envs/prod/events/" + eventID + "/votes/" + userID
}

func (paths *FirestorePaths) ObserversForEventCol(eventID string) string {
	return "envs/prod/event_observers/" + eventID + "/uids"
}

func (paths *FirestorePaths) ObserversForEventForUserDoc(eventID string, userID string) string {
	return "envs/prod/event_observers/" + eventID + "/uids/" + userID
}


func (paths *FirestorePaths) MessengerRecipientForUserDoc(userID string) string {
	return "envs/prod/messenger_recipients/" + userID
}
