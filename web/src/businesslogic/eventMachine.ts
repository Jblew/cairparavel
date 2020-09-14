
import { assign, Interpreter, Machine, send } from 'xstate'
import { Event } from './Event'

interface Schema {
  states: {
    InitialFetch: {},
    Error: {}
    TimeVoting: {}
    DoTimeVote: {}
    DoTimeUnvote: {}
    WaitingForTimeConfirm: {}
    DoTimeConfirm: {}
    Cancelled: {}
    MembersSignup: {}
    DoMemberSignup: {}
    DoMemberSignout: {}
    SignupClosed: {}
    InProggress: {}
    Finished: {}
    DoUpdateDetails: {}
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
  | { type: 'CONFIRM_TIME', startTime: number, endTime: string }
  | { type: 'SIGNUP_MEMBER' }
  | { type: 'SIGNOUT_MEMBER' }
  | { type: 'UPDATE_DETAILS', name: string, description: string }
  | { type: 'ERROR', message: string }

interface Context {
  currentUid: string
  eventId: string
  event?: Event
}

export const eventMachine = Machine<
  Context,
  Schema,
  Events
>(
  {
    id: 'event',
    initial: 'InitialFetch',
    invoke: { src: 'syncEvent' },
    on: {
      UPDATED: { actions: assign({ event: (_, evt: any) => evt.event }), },
      SYNC_ERROR: 'Error',
    },
    states: {
      InitialFetch: {
        on: {
          UPDATED: 'TimeVoting',
        }
      },
      TimeVoting: {
        on: {
          '': { target: 'WaitingForTimeConfirm', cond: (ctx) => now() >= ctx.event!.votingTime },
          TIME_VOTE: 'DoTimeVote',
          TIME_UNVOTE: 'DoTimeUnVote',
          UPDATE_DETAILS: { target: 'DoUpdateDetails', cond: (ctx) => ctx.currentUid === ctx.event!.ownerUid },
          UPDATED: '',
        },
        after: { 1000: '' }
      },
      DoTimeVote: {
        invoke: {
          src: 'timeVote',
          onDone: 'TimeVoting',
          onError: { target: 'TimeVoting', actions: ['logError', send({ type: 'ERROR', message: 'Failed to time vote' })] }
        }
      },
      DoTimeUnvote: {
        invoke: {
          src: 'timeUnvote',
          onDone: 'TimeVoting',
          onError: { target: 'TimeVoting', actions: ['logError', send({ type: 'ERROR', message: 'Failed to time unvote' })] }
        }
      },
      WaitingForTimeConfirm: {
        on: {
          '': [
            { target: 'MembersSignup', cond: (ctx) => ctx.event!.timeConfirmed },
            { target: 'TimeVoting', cond: (ctx) => now() < ctx.event!.votingTime },
            { target: 'Cancelled', cond: (ctx) => now() >= ctx.event!.startTime },
          ],
          UPDATE_DETAILS: { target: 'DoUpdateDetails', cond: (ctx) => ctx.currentUid === ctx.event!.ownerUid },
          CONFIRM_TIME: { target: 'DoTimeConfirm', cond: (ctx) => ctx.currentUid === ctx.event!.ownerUid && now() < ctx.event!.startTime },
          UPDATED: '',
        },
        after: { 1000: '' }
      },
      DoTimeConfirm: {
        invoke: {
          src: 'timeConfirm',
          onDone: 'MembersSignup',
          onError: { target: 'WaitingForTimeConfirm', actions: ['logError', send({ type: 'ERROR', message: 'Failed to confirm time' })] }
        }
      },
      MembersSignup: {
        on: {
          '': [
            { target: 'SignupClosed', cond: (ctx) => now() >= ctx.event!.signupTime && Object.keys(ctx.event!.signedMembers).length >= ctx.event!.minParticipants },
            { target: 'Cancelled', cond: (ctx) => now() >= ctx.event!.signupTime && Object.keys(ctx.event!.signedMembers).length < ctx.event!.minParticipants },
          ],
          SIGNUP_MEMBER: { target: 'DoMemberSignup', cond: (ctx) => Object.keys(ctx.event!.signedMembers).length < ctx.event!.maxParticipants && now() < ctx.event!.signupTime },
          SIGNOUT_MEMBER: { target: 'DoMemberSignout', cond: (ctx) => now() < ctx.event!.signupTime },
          UPDATED: '',
        },
        after: { 1000: '' }
      },
      DoMemberSignup: {
        invoke: {
          src: 'memberSignup',
          onDone: 'MembersSignup',
          onError: { target: 'MembersSignup', actions: ['logError', send({ type: 'ERROR', message: 'Cannot sign up' })] }
        }
      },
      DoMemberSignout: {
        invoke: {
          src: 'memberSignout',
          onDone: 'MembersSignup',
          onError: { target: 'MembersSignup', actions: ['logError', send({ type: 'ERROR', message: 'Cannot sign out' })] }
        }
      },
      SignupClosed: {
        on: {
          '': [
            { target: 'InProggress', cond: (ctx) => now() >= ctx.event!.startTime },
          ],
          UPDATED: '',
        },
        after: { 1000: '' }
      },
      InProggress: {
        on: {
          '': [
            { target: 'Finished', cond: (ctx) => now() >= ctx.event!.endTime },
          ],
          UPDATED: '',
        },
        after: { 1000: '' }
      },
      Finished: {},
      Cancelled: {},
      DoUpdateDetails: {
        invoke: {
          src: 'updateDetails',
          onDone: 'TimeVoting',
          onError: { target: 'TimeVoting', actions: ['logError', send({ type: 'ERROR', message: 'Cannot update details' })] }
        }
      },
      Error: {
        on: {
          UPDATED: 'TimeVoting',
        }
      }
    }
  },
  {
    actions: {
    },
  },
)

function now() {
  return Date.now() / 1000
}

export type EventMachineInterpreter = Interpreter<
  Context,
  Schema,
  Events
>

export type EventMachineContext = Context
