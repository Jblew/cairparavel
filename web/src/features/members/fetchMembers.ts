import firebase from 'firebase/app';
import { projectConfig } from '@/config';

export async function fetchMembers(): Promise<UserDoc[]> {
  const snapshot = await firebase.firestore()
    .collection(projectConfig.users.firestoreCollection)
    .orderBy('JoinedAt', 'desc')
    .get()
  if (snapshot.empty) return []
  return snapshot.docs.map(doc => doc.data() as UserDoc)
}

export interface UserDoc {
  Email: string
  UID: string
  JoinedAt: any
}
