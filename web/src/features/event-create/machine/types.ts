import { Event } from '@/businesslogic'

export interface Schema {
  states: {
    DetailsSetup: {}
    QuestionVoting: {
      states: {
        Initial: {}
        Voting: {},
        NoVoting: {}
      }
    }
    VotingSetup: {}
    EventTimeSetup: {}
    SignupSetup: {}
    ParticipantLimits: {}
    Confirm: {}
    DoSave: {}
    Success: {}
    Error: {}
  }
}


export type Events =
  | { type: 'BACK' }
  | { type: 'NEXT' }
  | {
    type: 'SET_DETAILS'
    details: { name: string, description: string }
  }
  | { type: 'CHOOSE_VOTING' }
  | { type: 'CHOOSE_NO_VOTING' }
  | {
    type: 'SET_VOTING'
    voting: {
      canSuggestTime: boolean,
      allowedTimes: number[]
      votingTime: number
    }
  }
  | {
    type: 'SET_EVENT_TIME'
    time: { startTime: number, endTime: number }
  }
  | {
    type: 'SET_SIGNUP_TIME'
    time: number
  }
  | {
    type: 'SET_PARTICIPANT_LIMITS'
    limits: { minParticipants: number, maxParticipants: number }
  }
  | {
    type: 'DO_SAVE'
  }

export interface Context {
  event: Event
}
