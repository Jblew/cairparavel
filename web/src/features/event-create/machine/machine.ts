
import { assign, Interpreter, Machine, send } from 'xstate'
import { Context, Events, Schema } from './types'


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
      assignDetails: assign({
        event: (ctx, evt: any) => ({ ...ctx.event, ...evt.details })
      }),
      enableVoting: assign({
        event: (ctx) => ({ ...ctx.event, votingTime: Date.now() + 14 * 24 * 3600 })
      }),
      disableVoting: assign({
        event: (ctx) => ({ ...ctx.event, votingTime: -1 })
      }),
      assignVoting: assign({
        event: (ctx, evt: any) => ({ ...ctx.event, ...evt.voting })
      }),
      assignEventTime: assign({
        event: (ctx, evt: any) => ({ ...ctx.event, ...evt.time })
      }),
      assignSignupTime: assign({
        event: (ctx, evt: any) => ({ ...ctx.event, signupTime: evt.time })
      }),
      assignParticipantLimits: assign({
        event: (ctx, evt: any) => ({ ...ctx.event, ...evt.limits })
      }),
    },
    guards: {
      isDetailsValid: (_, evt: any) => evt.details.name.length > 5 && evt.details.description.length > 5,
      isVotingValid: (_, evt: any) => evt.voting.votingTime > Date.now() && evt.voting.allowedTimes.every((time: number) => time > evt.voting.votingTime),
      isEventTimeValid: (_, evt: any) => evt.time.startTime < evt.time.endTime && evt.time.startTime > Date.now(),
      isSignupTimeValid: (_, evt: any) => evt.time > Date.now() && evt.time < evt.time.startTime,
      isParticipantLimitsValid: (_, evt: any) => evt.limits.minParticipants > 0 && evt.limits.maxParticipants > evt.limits.minParticipants
    }
  }
)

export type EventMachineInterpreter = Interpreter<
  Context,
  Schema,
  Events
>

export type EventMachineContext = Context
