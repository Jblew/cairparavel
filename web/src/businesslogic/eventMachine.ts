
import { assign, Interpreter, Machine, send } from 'xstate'
import { Event } from './Event'

interface Schema {
  states: {
    Error: {}
    TimeVoting: {}
    DoTimeVote: {}
    WaitingForTimeConfirm: {}
    DoTimeConfirm: {}
    Cancelled: {}
    MembersSignup: {}
    DoMemberSignup: {}
    DoMemberSignout: {}
    SignupClosed: {}
    InProggress: {}
    Finished: {}
  }
}

type Events =
  | {
    type: 'UPDATED'
    data: Event
  }
  | {
    type: 'TIME_VOTE'
    time: number
  }
  | { type: 'CONFIRM_TIME' }
  | { type: 'SIGNUP_MEMBER' }
  | { type: 'SIGNOUT_MEMBER' }
  | { type: 'ERROR', message: string }

interface Context {
  currentUid: string
  event: Event
}

export const eventMachine = Machine<
  Context,
  Schema,
  Events
>(
  {
    id: 'event',
    initial: 'TimeVoting',
    states: {
      TimeVoting: {
        on: {
          '': { target: 'WaitingForTimeConfirm', cond: (ctx) => nowS() >= ctx.event.votingTimeS },
          TIME_VOTE: 'DoTimeVote',
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
      WaitingForTimeConfirm: {
        on: {
          '': [
            { target: 'MembersSignup', cond: (ctx) => ctx.event.timeConfirmed },
            { target: 'TimeVoting', cond: (ctx) => nowS() < ctx.event.votingTimeS },
            { target: 'Cancelled', cond: (ctx) => nowS() >= ctx.event.startTimeS },
          ],
          CONFIRM_TIME: { target: 'DoTimeConfirm', cond: (ctx) => ctx.currentUid === ctx.event.ownerUid && nowS() < ctx.event.startTimeS },
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
            { target: 'SignupClosed', cond: (ctx) => nowS() >= ctx.event.signupTimeS && Object.keys(ctx.event.signedMembers).length >= ctx.event.minParticipants },
            { target: 'Cancelled', cond: (ctx) => nowS() >= ctx.event.signupTimeS && Object.keys(ctx.event.signedMembers).length < ctx.event.minParticipants },
          ],
          SIGNUP_MEMBER: { target: 'DoMemberSignup', cond: (ctx) => Object.keys(ctx.event.signedMembers).length < ctx.event.maxParticipants && nowS() < ctx.event.signupTimeS },
          SIGNOUT_MEMBER: { target: 'DoMemberSignout', cond: (ctx) => nowS() < ctx.event.signupTimeS },
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
            { target: 'InProggress', cond: (ctx) => nowS() >= ctx.event.startTimeS },
          ],
          UPDATED: '',
        },
        after: { 1000: '' }
      },
      InProggress: {
        on: {
          '': [
            { target: 'Finished', cond: (ctx) => nowS() >= ctx.event.endTimeS },
          ],
          UPDATED: '',
        },
        after: { 1000: '' }
      },
      Finished: {},
      Cancelled: {},
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

function nowS() {
  return Date.now() / 1000
}

export type EventMachineInterpreter = Interpreter<
  Context,
  Schema,
  Events
>
