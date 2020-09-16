import { isE2ETest } from '@/util'
import firebase from 'firebase/app'

const env = isE2ETest() ? 'test' : 'prod'
const envCol = `envs/${env}`

export const projectConfig = {
  firebaseAuth: {
    signInOptions: [firebase.auth.GoogleAuthProvider.PROVIDER_ID],
  },
  users: {
    firestoreCollection: `${envCol}/users`,
  },
  roles: {
    list: ['leader', 'member'],
    firestoreCollection: (roleName: string) => `${envCol}/roles/${roleName}/uids`,
  },
  events: {
    firestoreEventDoc: (eventId: string) => `${envCol}/events/${eventId}`,
    firestoreEventVoteDoc: (eventId: string, uid: string) =>
      `${envCol}/events/${eventId}/votes/${uid}`,
    firestoreEventSignupDoc: (eventId: string, uid: string) =>
      `${envCol}/events/${eventId}/signedMembers/${uid}`,
  },
}
