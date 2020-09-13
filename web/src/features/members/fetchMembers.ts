import firebase from 'firebase/app';

const usersColName = "users"

export async function fetchMembers(): Promise<UserDoc[]> {
  const snapshot = await firebase.firestore().collection(usersColName).orderBy('JoinedAt', 'desc').get()
  console.log(snapshot)
  if (snapshot.empty) return []
  return snapshot.docs.map(doc => doc.data() as UserDoc)
}

export interface UserDoc {
  Email: string
  UID: string
  JoinedAt: any
}
