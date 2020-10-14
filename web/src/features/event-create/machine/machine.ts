import { assign, Interpreter, Machine, send } from 'xstate'
import { Context, Events, Schema } from './types'
import { DateTime } from 'luxon'
import { validateEvent } from '@/businesslogic'

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
    canSuggestTime: true,
  },
}

export const createEventMachine = Machine<Context, Schema, Events>(
  {
    id: 'createEvent',
    initial: 'DetailsSetup',
    context: initialContext,
    states: {
      Initial: {
        on: {
          '': 'DetailsSetup',
        }
      },
      DetailsSetup: {
        entry: 'assignDefaultName',
        on: {
          SET_DETAILS: {
            target: 'QuestionVoting',
            cond: 'isDetailsValid',
            actions: ['assignDetails', 'validateEvent']
          },
        },
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
                { target: 'NoVoting' },
              ],
            },
          },
          Voting: {
            on: {
              NEXT: {
                target: '#createEvent.VotingSetup',
                actions: ['enableVoting', 'validateEvent'],
              },
            },
          },
          NoVoting: {
            on: {
              NEXT: {
                target: '#createEvent.EventTimeSetup',
                actions: ['disableVoting', 'validateEvent'],
              },
            },
          },
        },
      },
      VotingSetup: {
        entry: 'assignDefaultVoting',
        on: {
          SET_VOTING: {
            target: 'SignupSetup',
            cond: 'isVotingValid',
            actions: ['assignVoting', 'validateEvent'],
          },
          BACK: 'QuestionVoting',
        },
      },
      EventTimeSetup: {
        entry: 'assignDefaultEventTime',
        on: {
          SET_EVENT_TIME: {
            target: 'SignupSetup',
            cond: 'isEventTimeValid',
            actions: ['assignEventTime', 'validateEvent'],
          },
          BACK: 'QuestionVoting',
        },
      },
      SignupSetup: {
        on: {
          SET_SIGNUP_TIME: {
            target: 'ParticipantLimits',
            cond: 'isSignupTimeValid',
            actions: ['assignSignupTime', 'validateEvent'],
          },
          BACK: 'QuestionVoting',
        },
      },
      ParticipantLimits: {
        entry: 'assignDefaultParticipantLimits',
        on: {
          SET_PARTICIPANT_LIMITS: {
            target: 'Confirm',
            cond: 'isParticipantLimitsValid',
            actions: ['assignParticipantLimits', 'validateEvent'],
          },
          BACK: 'SignupSetup',
        },
      },
      Confirm: {
        on: {
          DO_SAVE: 'DoSave',
          BACK: 'SignupSetup',
        },
      },
      DoSave: {
        invoke: {
          src: 'saveEvent',
          onDone: 'Success',
          onError: { target: 'Error', actions: 'logError' },
        },
      },
      Success: {
        on: {
          BACK: 'Confirm',
        },
      },
      Error: {
        on: {
          BACK: 'Confirm',
        },
      },
    },
  },
  {
    actions: {
      logError: (_, { data }: any) => console.error(new Error(data)),
      validateEvent: ({ event }) => validateEvent(event),
      assignDetails: assign({
        event: (ctx, evt: any) => ({ ...ctx.event, ...evt.details }),
      }),
      enableVoting: assign({
        event: ctx => ({
          ...ctx.event,
          votingTime: Date.now() + 14 * 24 * 3600,
        }),
      }),
      disableVoting: assign({
        event: ctx => ({ ...ctx.event, votingTime: -1, timeConfirmed: true }),
      }),
      assignVoting: assign({
        event: (ctx, evt: any) => ({
          ...ctx.event,
          ...evt.voting,
          startTime: DateTime.fromMillis(evt.voting.votingTime).plus({ hours: 24 }).toMillis(),
          endTime: DateTime.fromMillis(evt.voting.votingTime).plus({ hours: 25 }).toMillis(),
        }),
      }),
      assignEventTime: assign({
        event: (ctx, evt: any) => {
          return ({ ...ctx.event, startTime: Number(evt.time.startTime), endTime: Number(evt.time.endTime) })
        },
      }),
      assignSignupTime: assign({
        event: (ctx, evt: any) => ({ ...ctx.event, signupTime: Number(evt.time) }),
      }),
      assignParticipantLimits: assign({
        event: (ctx, evt: any) => ({ ...ctx.event, ...evt.limits }),
      }),
      assignDefaultName: assign({
        event: (ctx) => ({
          ...ctx.event,
          name: 'Event #' + Math.random(),
          description: 'Description of your event...',
        })
      }),
      assignDefaultVoting: assign({
        event: (ctx) => ({
          ...ctx.event,
          votingTime: DateTime.local().plus({ days: 5 }).toMillis(),
          signupTime: DateTime.local().plus({ days: 6 }).toMillis(),
          allowedTimes: [
            DateTime.local().plus({ days: 7 }).toMillis(),
            DateTime.local().plus({ days: 8 }).toMillis()
          ]
        })
      }),
      assignDefaultEventTime: assign({
        event: (ctx) => ({
          ...ctx.event,
          signupTime: DateTime.local().plus({ days: 6 }).toMillis(),
          startTime: DateTime.local().plus({ days: 7 }).toMillis(),
          endTime: DateTime.local().plus({ days: 7 }).plus({ hours: 2 }).toMillis(),
        })
      }),
      assignDefaultParticipantLimits: assign({
        event: (ctx) => ({
          ...ctx.event,
          minParticipants: 1,
          maxParticipants: 100,
        })
      }),
    },
    guards: {
      isDetailsValid: (_, evt: any) =>
        evt.details.name.length > 5 && evt.details.description.length > 5,
      isVotingValid: (_, evt: any) =>
        evt.voting.votingTime > Date.now() &&
        evt.voting.allowedTimes.every(
          (time: number) => time > evt.voting.votingTime,
        ),
      isEventTimeValid: (_, evt: any) =>
        evt.time.startTime < evt.time.endTime &&
        evt.time.startTime > Date.now(),
      isSignupTimeValid: (_, evt: any) =>
        evt.time > Date.now(),
      isParticipantLimitsValid: (_, evt: any) =>
        evt.limits.minParticipants > 0 &&
        evt.limits.maxParticipants > evt.limits.minParticipants,
    },
  },
)

export type CreateEventInterpreter = Interpreter<Context, Schema, Events>

export type CreateEventContext = Context
