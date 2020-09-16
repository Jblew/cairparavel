import { Event } from '@/businesslogic';

import { projectConfig } from '@/config'
import firebase from 'firebase/app'

export function saveEvent(event: Event) {
  const colPath = projectConfig.events.firestoreEventDoc('')
  const eventOwned: Event = {
    ...event,
    ownerUid: firebase.auth().currentUser!.uid
  }
  delete eventOwned['votes']
  delete eventOwned['signedMembers']

  return firebase.firestore().collection(colPath).add(eventOwned)
}
