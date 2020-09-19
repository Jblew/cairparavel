export interface Event {
  id?: string
  ownerUid: Uid
  ownerDisplayName: UserDisplayName
  votingTime: number
  startTime: number
  endTime: number
  timeConfirmed: boolean
  signupTime: number
  votes?: Record<Uid, EventTimeVotes>
  signedMembers?: Record<Uid, EventSignup>
  minParticipants: number
  maxParticipants: number
  name: string
  description: string
  allowedTimes: number[]
  canSuggestTime: boolean
}

export interface EventTimeVotes {
  displayName: UserDisplayName
  times: number[]
}

export interface EventSignup {
  displayName: UserDisplayName
}

type Uid = string
type UserDisplayName = string
