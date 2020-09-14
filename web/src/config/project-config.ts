import firebase from 'firebase/app';

export const projectConfig = {
  firebaseAuth: {
    signInOptions: [firebase.auth.GoogleAuthProvider.PROVIDER_ID],
  },
  users: {
    firestoreCollection: 'users',
  },
  roles: {
    list: ['leader', 'member'],
    firestoreCollection: (roleName: string) => `roles/${roleName}/uids`
  },
  events: {
    firestoreEventDoc: (eventId: string) => `events/${eventId}`,
    firestoreEventVoteDoc: (eventId: string, uid: string) => `events/${eventId}/votes/${uid}`,
    firestoreEventSignupDoc: (eventId: string, uid: string) => `events/${eventId}/signedMembers/${uid}`
  },
};
