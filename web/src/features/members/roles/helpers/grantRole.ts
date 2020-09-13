import firebase from 'firebase/app';
import { projectConfig } from '@/config';

const recordPlaceholder = { a: true }

export async function grantRole(uid: string, role: string) {
  await firebase.firestore()
    .collection(projectConfig.roles.firestoreCollection(role))
    .doc(uid).set(recordPlaceholder)
}
