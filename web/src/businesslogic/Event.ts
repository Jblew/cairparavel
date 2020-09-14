
export interface Event {
  ownerUid: Uid
  ownerDisplayName: UserDisplayName
  votingTimeS: number
  startTimeS: number
  endTimeS: number
  timeConfirmed: boolean
  signupTimeS: number,
  votes: EventTimeVote[]
  signedMembers: Record<Uid, UserDisplayName>
  minParticipants: number
  maxParticipants: number
}

interface EventTimeVote {
  uid: Uid,
  userDisplayName: UserDisplayName
  timeS: number
}

type Uid = string
type UserDisplayName = string
