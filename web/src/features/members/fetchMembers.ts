import firebase from 'firebase/app';

const usersColName = "users"

export async function fetchMembers(): Promise<UserDoc[]> {
  const snapshot = await firebase.firestore().collection(usersColName).orderBy('joinedAt', 'desc').get()
  if (snapshot.empty) return []
  return snapshot.docs.map(doc => doc.data() as UserDoc)
}

export interface UserDoc {
  email: string
  uid: string
  joinedAt: any
}
