import { interpret } from 'xstate'
import {
  eventMachineFactory,
  EventMachineInterpreter,
  Event,
} from '@/businesslogic'
import firebase from 'firebase/app'
import {
  syncEventActorFactory,
  timeVote,
  timeUnvote,
  timeConfirm,
  memberSignup,
  memberSignout,
  updateDetails,
} from './services'
import Vue from 'vue'

export function getEventInterpreter({ eventId, event }: { eventId: string, event?: Event }): EventMachineInterpreter {
  const eventMachine = eventMachineFactory({
    now,
    syncActor: syncEventActorFactory(eventId),
  })
  const currentUid = firebase.auth().currentUser?.uid
  if (!currentUid)
    throw new Error(
      'Firebase currentUid must be populated before starting EventInterpreter',
    )
  const userDisplayName = firebase.auth().currentUser!.displayName || ''
  const context = Vue.observable({ currentUid, userDisplayName, eventId, event, })

  return interpret(
    eventMachine.withContext(context).withConfig({
      services: {
        timeVote,
        timeUnvote,
        timeConfirm,
        memberSignup,
        memberSignout,
        updateDetails,
      },
    }),
  )
}

function now() {
  return Date.now()
}
