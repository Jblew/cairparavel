import { isE2ETest, toPrettySlug } from '@/util'
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
    firestoreEventCol: `${envCol}/events/`,
    firestoreEventVoteDoc: (eventId: string, uid: string) =>
      `${envCol}/events/${eventId}/votes/${uid}`,
    firestoreEventSignupDoc: (eventId: string, uid: string) =>
      `${envCol}/events/${eventId}/signedMembers/${uid}`,
    eventUrl: ({ name, id }: { name: string, id?: string }) => `/event/${toPrettySlug(name)}/${id}`
  },
}
