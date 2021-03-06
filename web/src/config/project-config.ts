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
  messengerChatUrl: (firebaseUid: string) => `https://m.me/103901188157799?ref=${firebaseUid}`,
  events: {
    firestoreEventCol: `${envCol}/events/`,
    firestoreEventVoteCol: (eventId: string) =>
      `${envCol}/events/${eventId}/votes`,
    firestoreEventVoteDoc: (eventId: string, uid: string) =>
      `${envCol}/events/${eventId}/votes/${uid}`,
    firestoreEventSignupCol: (eventId: string) =>
      `${envCol}/events/${eventId}/signedMembers`,
    firestoreEventSignupDoc: (eventId: string, uid: string) =>
      `${envCol}/events/${eventId}/signedMembers/${uid}`,
    observerDoc: (eventId: string, uid: string) =>
      `${envCol}/event_observers/${eventId}/uids/${uid}`,
    eventUrl: ({ name, id }: { name: string, id?: string }) => `/event/${toPrettySlug(name)}/${id}`,
  },
  comments: {
    firestoreEventCommentsCol: (eventId: string) => `${envCol}/event_comments/${eventId}/messages`,
  }
}
