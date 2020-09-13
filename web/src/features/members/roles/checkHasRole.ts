import firebase from 'firebase/app';
import { projectConfig } from '@/config';

export async function checkHasRole(uid: string, role: string): Promise<boolean> {
  const snapshot = await firebase.firestore()
    .collection(projectConfig.roles.firestoreCollection(role))
    .doc(uid).get()

  return snapshot.exists
}
