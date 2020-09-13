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
};
