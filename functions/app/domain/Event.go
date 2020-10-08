package domain

type Event struct {
	ID               string                    `json:"id"`
	OwnerUid         string                    `json:"ownerUid"`
	OwnerDisplayName string                    `json:"ownerDisplayName"`
	VotingTime       int64                     `json:"votingTime"`
	StartTime        int64                     `json:"startTime"`
	EndTime          int64                     `json:"endTime"`
	TimeConfirmed    bool                      `json:"timeConfirmed"`
	SignupTime       int64                     `json:"signupTime"`
	Votes            map[string]EventTimeVotes `json:"votes"`
	SignedMembers    map[string]EventSignup    `json:"signedMembers"`
	MinParticipants  int                       `json:"minParticipants"`
	MaxParticipants  int                       `json:"maxParticipants"`
	Name             string                    `json:"name"`
	Description      string                    `json:"description"`
	AllowedTimes     []int64                   `json:"allowedTimes"`
	CanSuggestTime   bool                      `json:"canSuggestTime"`
}

type EventTimeVotes struct {
	UID         string  `json:"uid"`
	DisplayName string  `json:"displayName"`
	Times       []int64 `json:"times"`
}

type EventSignup struct {
	UID         string `json:"uid"`
	DisplayName string `json:"displayName"`
}
