
import { assign, Interpreter, Machine, send } from 'xstate'
import { Event } from '@/businesslogic'

const every = (...guards: any[]) => ({
  type: 'every',
  guards
});

interface Schema {
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

type Events =
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

interface Context {
  event: Event
}

const initialContext: Context = {
  event: {
    ownerUid: '',
    ownerDisplayName: '',
    votingTime: 0,
    startTime: 0,
    endTime: 0,
    timeConfirmed: false,
    signupTime: 0,
    votes: {},
    signedMembers: {},
    minParticipants: 0,
    maxParticipants: 0,
    name: '',
    description: '',
    allowedTimes: [],
    canSuggestTime: true
  }
}

export const createEventMachine = Machine<
  Context,
  Schema,
  Events
>(
  {
    id: 'creaeEvent',
    initial: 'DetailsSetup',
    context: initialContext,
    states: {
      DetailsSetup: {
        on: {
          SET_DETAILS: { target: 'QuestionVoting', cond: 'isDetailsValid', actions: 'assignDetails' },
        }
      },
      QuestionVoting: {
        on: {
          CHOOSE_VOTING: '.Voting',
          CHOOSE_NO_VOTING: '.NoVoting',
          BACK: 'DetailsSetup',
        },
        initial: 'Initial',
        states: {
          Initial: {
            on: {
              '': [
                { target: 'Voting', cond: ctx => ctx.event!.votingTime != 0 },
                { target: 'NoVoting' }
              ]
            }
          },
          Voting: {
            on: {
              NEXT: { target: 'VotingSetup', actions: 'enableVoting' }
            }
          },
          NoVoting: {
            on: {
              NEXT: { target: 'EventTimeSetup', actions: 'disableVoting' }
            }
          }
        }
      },
      VotingSetup: {
        on: {
          SET_VOTING: { target: 'SignupSetup', cond: 'isVotingValid', actions: 'assignVoting' },
          BACK: 'QuestionVoting',
        }
      },
      EventTimeSetup: {
        on: {
          SET_EVENT_TIME: { target: 'SignupSetup', cond: 'isEventTimeValid', actions: 'assignEventTime' },
          BACK: 'QuestionVoting',
        }
      },
      SignupSetup: {
        on: {
          SET_SIGNUP_TIME: { target: 'ParticipantLimits', cond: 'isSignupTimeValid', actions: 'assignSignupTime' },
          BACK: 'QuestionVoting',
        }
      },
      ParticipantLimits: {
        on: {
          SET_PARTICIPANT_LIMITS: { target: 'Confirm', cond: 'isParticipantLimitsValid', actions: 'assignParticipantLimits' },
          BACK: 'SignupSetup',
        }
      },
      Confirm: {
        on: {
          DO_SAVE: 'DoSave',
          BACK: 'SignupSetup',
        }
      },
      DoSave: {
        invoke: {
          src: 'saveEvent',
          onDone: 'Success',
          onError: { target: 'Error', actions: 'logError' }
        }
      },
      Success: {
        on: {
          BACK: 'Confirm'
        }
      },
      Error: {
        on: {
          BACK: 'Confirm'
        }
      }
    },
  },
  {
    actions: {
      logError: (_, { data }: any) => console.error(new Error(data)),
    },
    guards: {
      every: (ctx, event, { guard }: any) => {
        const { guards } = guard;
        return guards.every((guardKey: any) => guards[guardKey](ctx, event));
      }
    }
  }
)


export type EventMachineInterpreter = Interpreter<
  Context,
  Schema,
  Events
>

export type EventMachineContext = Context
