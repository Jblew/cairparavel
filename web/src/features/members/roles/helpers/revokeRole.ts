import firebase from 'firebase/app'
import { projectConfig } from '@/config'

export async function revokeRole(uid: string, role: string) {
  await firebase
    .firestore()
    .collection(projectConfig.roles.firestoreCollection(role))
    .doc(uid)
    .delete()
}
