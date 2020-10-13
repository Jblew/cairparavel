import { Actor, assign, Interpreter, InvokeCallback, Machine, send, spawn, Subscribable } from 'xstate'
import { Event } from './Event'

interface Schema {
  states: {
    InitialFetch: {}
    Error: {}
    TimeVoting: {}
    DoTimeVote: {}
    DoTimeUnvote: {}
    WaitingForTimeConfirm: {}
    DoTimeConfirm: {}
    Cancelled: {}
    MembersSignup: {
      states: {
        Initial: {}
        SignedUp: {}
        SignedOut: {}
      }
    },
    DoMemberSignup: {}
    DoMemberSignout: {}
    SignupClosed: {}
    InProggress: {}
    Finished: {}
    DoUpdateDetails: {}
    DoDelete: {}
    Deleted: {}
  }
}

type Events =
  | {
    type: 'UPDATED'
    data: Event
  }
  | {
    type: 'SYNC_ERROR'
  }
  | {
    type: 'TIME_VOTE'
    time: number
  }
  | {
    type: 'TIME_UNVOTE'
    time: number
  }
  | { type: 'CONFIRM_TIME'; startTime: number; endTime: string }
  | { type: 'SIGNUP_MEMBER' }
  | { type: 'SIGNOUT_MEMBER' }
  | { type: 'UPDATE_DETAILS'; name: string; description: string }
  | { type: 'ERROR'; message: string }
  | { type: 'DELETE' }

interface Context {
  currentUid: string
  userDisplayName: string
  eventId: string
  event?: Event
  syncActorRef?: Actor<any>
}

export function eventMachineFactory({ now, syncActor }: { now(): number, syncActor: InvokeCallback }) {
  return Machine<Context, Schema, Events>(
    {
      id: 'event',
      initial: 'InitialFetch',
      on: {
        UPDATED: { actions: 'assignUpdatedEvent', },
        SYNC_ERROR: 'Error',
      },
      states: {
        InitialFetch: {
          entry: [assign<any, any>({
            syncActorRef: () => spawn(syncActor)
          })],
          on: {
            UPDATED: { target: 'TimeVoting', actions: 'assignUpdatedEvent' },
          },
        },
        TimeVoting: {
          on: {
            '': {
              target: 'WaitingForTimeConfirm',
              cond: ctx => now() >= ctx.event!.votingTime,
            },
            TIME_VOTE: 'DoTimeVote',
            TIME_UNVOTE: 'DoTimeUnvote',
            UPDATE_DETAILS: {
              target: 'DoUpdateDetails',
              cond: ctx => ctx.currentUid === ctx.event!.ownerUid,
            },
            DELETE: {
              target: 'DoDelete',
              cond: ctx => ctx.currentUid === ctx.event!.ownerUid,
            },
            UPDATED: { actions: 'assignUpdatedEvent' },
          },
          after: { 1000: '' },
        },
        DoTimeVote: {
          invoke: {
            src: 'timeVote',
            onDone: 'TimeVoting',
            onError: {
              target: 'TimeVoting',
              actions: [
                'logError',
                send({ type: 'ERROR', message: 'Failed to time vote' }),
              ],
            },
          },
        },
        DoTimeUnvote: {
          invoke: {
            src: 'timeUnvote',
            onDone: 'TimeVoting',
            onError: {
              target: 'TimeVoting',
              actions: [
                'logError',
                send({ type: 'ERROR', message: 'Failed to time unvote' }),
              ],
            },
          },
        },
        WaitingForTimeConfirm: {
          on: {
            '': [
              {
                target: 'MembersSignup',
                cond: ctx => ctx.event!.timeConfirmed,
              },
              {
                target: 'TimeVoting',
                cond: ctx => now() < ctx.event!.votingTime,
              },
              {
                target: 'Cancelled',
                cond: ctx => now() >= ctx.event!.startTime,
              },
            ],
            UPDATE_DETAILS: {
              target: 'DoUpdateDetails',
              cond: ctx => ctx.currentUid === ctx.event!.ownerUid,
            },
            DELETE: {
              target: 'DoDelete',
              cond: ctx => ctx.currentUid === ctx.event!.ownerUid,
            },
            CONFIRM_TIME: {
              target: 'DoTimeConfirm',
              cond: ctx =>
                ctx.currentUid === ctx.event!.ownerUid &&
                now() < ctx.event!.startTime,
            },
            UPDATED: { actions: 'assignUpdatedEvent' },
          },
          after: { 1000: '' },
        },
        DoTimeConfirm: {
          invoke: {
            src: 'timeConfirm',
            onDone: 'MembersSignup',
            onError: {
              target: 'WaitingForTimeConfirm',
              actions: [
                'logError',
                send({ type: 'ERROR', message: 'Failed to confirm time' }),
              ],
            },
          },
        },
        MembersSignup: {
          initial: 'Initial',
          on: {
            '': [
              {
                target: 'SignupClosed',
                cond: ctx =>
                  now() >= ctx.event!.signupTime &&
                  Object.keys(ctx.event!.signedMembers || {}).length >=
                  ctx.event!.minParticipants,
              },
              {
                target: 'Cancelled',
                cond: ctx =>
                  now() >= ctx.event!.signupTime &&
                  Object.keys(ctx.event!.signedMembers || {}).length <
                  ctx.event!.minParticipants,
              },
            ],
            UPDATE_DETAILS: {
              target: 'DoUpdateDetails',
              cond: ctx => ctx.currentUid === ctx.event!.ownerUid,
            },
            DELETE: {
              target: 'DoDelete',
              cond: ctx => ctx.currentUid === ctx.event!.ownerUid,
            },
            UPDATED: { target: '.Initial', actions: 'assignUpdatedEvent' },
          },
          after: { 1000: '.Initial' },
          states: {
            Initial: {
              on: {
                '': [
                  {
                    target: 'SignedUp', cond: (ctx) => Object.keys(ctx.event!.signedMembers || {}).includes(ctx.currentUid)
                  },
                  { target: 'SignedOut' },
                ]
              }
            },
            SignedUp: {
              on: {
                SIGNOUT_MEMBER: {
                  target: '#event.DoMemberSignout',
                },
              }
            },
            SignedOut: {
              on: {
                SIGNUP_MEMBER: {
                  target: '#event.DoMemberSignup',
                },
              }
            }
          }
        },
        DoMemberSignup: {
          invoke: {
            src: 'memberSignup',
            onDone: 'MembersSignup.SignedUp',
            onError: {
              target: 'MembersSignup',
              actions: [
                'logError',
                send({ type: 'ERROR', message: 'Cannot sign up' }),
              ],
            },
          },
        },
        DoMemberSignout: {
          invoke: {
            src: 'memberSignout',
            onDone: 'MembersSignup.SignedOut',
            onError: {
              target: 'MembersSignup',
              actions: [
                'logError',
                send({ type: 'ERROR', message: 'Cannot sign out' }),
              ],
            },
          },
        },
        SignupClosed: {
          on: {
            '': [
              {
                target: 'InProggress',
                cond: ctx => now() >= ctx.event!.startTime,
              },
            ],
            UPDATED: { actions: 'assignUpdatedEvent' },
          },
          after: { 1000: '' },
        },
        InProggress: {
          on: {
            '': [
              { target: 'Finished', cond: ctx => now() >= ctx.event!.endTime },
            ],
            UPDATED: { actions: 'assignUpdatedEvent' },
          },
          after: { 1000: '' },
        },
        Finished: {},
        Cancelled: {},
        DoUpdateDetails: {
          invoke: {
            src: 'updateDetails',
            onDone: 'TimeVoting',
            onError: {
              target: 'TimeVoting',
              actions: [
                'logError',
                send({ type: 'ERROR', message: 'Cannot update details' }),
              ],
            },
          },
        },
        Error: {
          on: {
            UPDATED: { target: 'TimeVoting', actions: 'assignUpdatedEvent' },
          },
        },
        DoDelete: {
          invoke: {
            src: 'deleteEvent',
            onDone: 'Deleted',
            onError: {
              target: 'TimeVoting',
              actions: [
                'logError',
                send({ type: 'ERROR', message: 'Cannot delete event' }),
              ],
            },
          },
        },
        Deleted: {}
      },
    },
    {
      actions: {
        logError: (_, { data }: any) => console.error(new Error(data)),
        assignUpdatedEvent: assign({ event: (_, evt: any) => evt.event })
      },
    },
  )
}

export type EventMachineInterpreter = Interpreter<Context, Schema, Events>

export type EventMachineContext = Context
