package config

// FirestorePathsStruct is a list of paths or path returners for firestore
type FirestorePathsStruct struct {
}

// ProdUsersCol path
func (paths *FirestorePathsStruct) ProdUsersCol() string {
	return "envs/prod/users"
}

// ProdUserDoc path
func (paths *FirestorePathsStruct) ProdUserDoc(userID string) string {
	return "envs/prod/users/" + userID
}

// TestUsersCol path
func (paths *FirestorePathsStruct) TestUsersCol() string {
	return "envs/test/users"
}

// TestUserDoc path
func (paths *FirestorePathsStruct) TestUserDoc(userID string) string {
	return "envs/test/users/" + userID
}

// MessengerNotificationsForUserCol path
func (paths *FirestorePathsStruct) MessengerNotificationsForUserCol(uid string) string {
	return "envs/prod/notifications/" + uid + "/messenger_queue"
}

// NotificationsHistoryForUserCol path
func (paths *FirestorePathsStruct) NotificationsHistoryForUserCol(uid string) string {
	return "envs/prod/notifications/" + uid + "/history"
}

// CommentsForEventCol path
func (paths *FirestorePathsStruct) CommentsForEventCol(eventID string) string {
	return "envs/prod/event_comments/" + eventID + "/messages"
}

// EventsCol path
func (paths *FirestorePathsStruct) EventsCol() string {
	return "envs/prod/events"
}

// SignupsForEventCol path
func (paths *FirestorePathsStruct) SignupsForEventCol(eventID string) string {
	return "envs/prod/events/" + eventID + "/signedMembers"
}

// SignupsForEventForUserDoc path
func (paths *FirestorePathsStruct) SignupsForEventForUserDoc(eventID string, userID string) string {
	return "envs/prod/events/" + eventID + "/signedMembers/" + userID
}

// VotesForEventCol path
func (paths *FirestorePathsStruct) VotesForEventCol(eventID string) string {
	return "envs/prod/events/" + eventID + "/votes"
}

// VotesForEventForUserDoc path
func (paths *FirestorePathsStruct) VotesForEventForUserDoc(eventID string, userID string) string {
	return "envs/prod/events/" + eventID + "/votes/" + userID
}

// ObserversForEventCol path
func (paths *FirestorePathsStruct) ObserversForEventCol(eventID string) string {
	return "envs/prod/event_observers/" + eventID + "/uids"
}

// ObserversForEventForUserDoc path
func (paths *FirestorePathsStruct) ObserversForEventForUserDoc(eventID string, userID string) string {
	return "envs/prod/event_observers/" + eventID + "/uids/" + userID
}

// MessengerRecipientForUserDoc path
func (paths *FirestorePathsStruct) MessengerRecipientForUserDoc(userID string) string {
	return "envs/prod/messenger_recipients/" + userID
}

// FirestorePaths paths to firestore documents
var FirestorePaths FirestorePathsStruct = FirestorePathsStruct{}
