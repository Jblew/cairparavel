import firebase from 'firebase/app'
import { projectConfig } from '@/config'

export async function fetchEvents(): Promise<EventDoc[]> {
  const snapshot = await firebase
    .firestore()
    .collection(projectConfig.events.firestoreEventDoc(''))
    .orderBy('startTime', 'asc')
    .get()
  if (snapshot.empty) return []
  return snapshot.docs.map(doc => doc.data() as EventDoc)
}

export interface EventDoc {
  Id: string
  SortingDate: any
}
