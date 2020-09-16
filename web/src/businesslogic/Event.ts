export interface Event {
  ownerUid: Uid
  ownerDisplayName: UserDisplayName
  votingTime: number
  startTime: number
  endTime: number
  timeConfirmed: boolean
  signupTime: number
  votes: Record<Uid, EventTimeVotes>
  signedMembers: Record<Uid, { displayName: UserDisplayName }>
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

type Uid = string
type UserDisplayName = string
