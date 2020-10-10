import ow from 'ow'

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

export function validateEvent(e: Event) {
  ow(e, 'Event', ow.object)
  ow(e.ownerUid, 'Event.ownerUid', ow.string)
  ow(e.ownerDisplayName, 'Event.ownerDisplayName', ow.string)
  ow(e.votingTime, 'Event.votingTime', ow.number)
  ow(e.startTime, 'Event.startTime', ow.number)
  ow(e.endTime, 'Event.endTime', ow.number)
  ow(e.timeConfirmed, 'Event.timeConfirmed', ow.boolean)
  ow(e.signupTime, 'Event.signupTime', ow.number.greaterThanOrEqual(0).integer)
  ow(e.minParticipants, 'Event.minParticipants', ow.number)
  ow(e.maxParticipants, 'Event.maxParticipants', ow.number)
  ow(e.name, 'Event.name', ow.string)
  ow(e.description, 'Event.description', ow.string)
  ow(e.allowedTimes, 'Event.allowedTimes', ow.array.ofType(ow.number))
  ow(e.canSuggestTime, 'Event.canSuggestTime', ow.boolean)
}

export interface EventTimeVotes {
  uid?: string
  displayName: UserDisplayName
  times: number[]
}

export interface EventSignup {
  uid?: string
  displayName: UserDisplayName
}

type Uid = string
type UserDisplayName = string
