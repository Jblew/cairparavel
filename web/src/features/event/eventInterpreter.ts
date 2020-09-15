import { interpret } from 'xstate';
import { eventMachineFactory } from '@/businesslogic/eventMachine';
import firebase from 'firebase/app'
import { syncEventActorFactory, timeVote, timeUnvote, timeConfirm, memberSignup, memberSignout, updateDetails } from './services'

export function getEventInterpreter(eventId: string) {
  const eventMachine = eventMachineFactory({ now })
  const currentUid = firebase.auth().currentUser?.uid
  if (!currentUid) throw new Error('Firebase currentUid must be populated before starting EventInterpreter')
  return interpret(eventMachine.withContext({ currentUid, eventId }).withConfig({
    services: {
      syncEvent: syncEventActorFactory(eventId),
      timeVote,
      timeUnvote,
      timeConfirm,
      memberSignup,
      memberSignout,
      updateDetails,
    },
  }))
}

function now() {
  return Date.now()
}
