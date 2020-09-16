import firebase from 'firebase/app'
import { projectConfig } from '@/config'

export async function fetchEvents(): Promise<EventDoc[]> {
  const snapshot = await firebase
    .firestore()
    .collection(projectConfig.users.firestoreCollection)
    .orderBy('SortingDate', 'desc')
    .get()
  if (snapshot.empty) return []
  return snapshot.docs.map(doc => doc.data() as EventDoc)
}

export interface EventDoc {
  Id: string
  SortingDate: any
}
