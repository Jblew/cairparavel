import { Event } from '@/businesslogic';

import { projectConfig } from '@/config'
import firebase from 'firebase/app'

export function saveEvent(event: Event) {
  const colPath = projectConfig.events.firestoreEventDoc('')
  return firebase.firestore().collection(colPath).add(event)
}
