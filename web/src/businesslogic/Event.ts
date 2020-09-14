
export interface Event {
  ownerUid: Uid
  ownerDisplayName: UserDisplayName
  votingTime: number
  startTime: number
  endTime: number
  timeConfirmed: boolean
  signupTime: number,
  votes: Record<Uid, EventTimeVote>
  signedMembers: Record<Uid, UserDisplayName>
  minParticipants: number
  maxParticipants: number
  name: string
  description: string
  allowedTimes: number[]
  canSuggestTime: boolean
}

interface EventTimeVote {
  userDisplayName: UserDisplayName
  time: number
}

type Uid = string
type UserDisplayName = string
