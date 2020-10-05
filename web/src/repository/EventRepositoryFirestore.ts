import { projectConfig } from '@/config'
import firebase from 'firebase/app'
import { Event, EventRepository, EventSignup, EventTimeVotes } from '@/businesslogic'

export class EventRepositoryFirestore implements EventRepository {
  subscribeToEvent({ eventId, onUpdated, onError }: { eventId: string, onUpdated: (e: Event) => void, onError: (err: Error) => void }) {
    const unsubscribeFn =
      this.getEventDocRef(eventId)
        .onSnapshot(snapshot => onUpdated(snapshot.data()!), onError)
    return unsubscribeFn
  }

  async saveEvent(event: Event) {
    const colRef = this.getEventColRef()
    const eventOwned: Event = {
      ...event,
      ownerUid: firebase.auth().currentUser!.uid,
      ownerDisplayName: firebase.auth().currentUser!.displayName || ''
    }
    delete eventOwned['votes']
    delete eventOwned['signedMembers']

    const result = await colRef.add(eventOwned)
    return result.id
  }

  async fetchEvents(): Promise<Event[]> {
    const snapshot = await this.getEventColRef()
      .orderBy('startTime', 'desc')
      .get()
    if (snapshot.empty) return []
    return snapshot.docs.map(doc => doc.data())
  }


  async timeVote({ eventId, currentUid, time }: { eventId: string, currentUid: string, time: number }) {
    const docRef = this.getEventVoteDocRef({ eventId, currentUid })
    const displayName = firebase.auth().currentUser!.displayName || 'Unknown user'
    const currentVoteSnapshot = await docRef.get()
    const times: number[] = currentVoteSnapshot.exists
      ? [...currentVoteSnapshot.data()!.times, time]
      : [time]
    const votes: EventTimeVotes = {
      displayName,
      times,
    }
    return docRef.set(votes)
  }

  async timeUnvote({ eventId, currentUid, time }: { eventId: string, currentUid: string, time: number }) {
    const docRef = this.getEventVoteDocRef({ eventId, currentUid })
    const displayName = firebase.auth().currentUser!.displayName || 'Unknown'
    const currentVoteSnapshot = await docRef.get()
    const times: number[] = currentVoteSnapshot.exists
      ? [
        ...currentVoteSnapshot
          .data()!
          .times.filter((comparedTime: any) => comparedTime != time),
      ]
      : []
    const votes: EventTimeVotes = {
      displayName,
      times,
    }
    return docRef.set(votes)
  }

  async timeConfirm(
    { eventId, startTime, endTime }: { eventId: string, startTime: number; endTime: number }
  ) {
    const docRef = this.getEventDocRef(eventId)
    const update: Partial<Event> = {
      timeConfirmed: true,
      startTime,
      endTime,
    }
    return docRef.set(update as Event)
  }


  async memberSignup({ eventId, currentUid }: { eventId: string, currentUid: string }) {
    const docRef = this.getEventSignupDoc({ eventId, currentUid })
    const displayName = firebase.auth().currentUser!.displayName || 'Unknown'
    return docRef.set({ displayName })
  }

  async memberSignout({ eventId, currentUid }: { eventId: string, currentUid: string }) {
    const docRef = this.getEventSignupDoc({ eventId, currentUid })
    return docRef.delete()
  }

  async updateDetails(
    { eventId, name, description }: { eventId: string, name: string; description: string },
  ) {
    const docRef = this.getEventDocRef(eventId)
    return docRef.set({
      name,
      description,
    } as Event)
  }

  private getEventColRef() {
    const docPath = projectConfig.events.firestoreEventCol
    return firebase
      .firestore()
      .collection(docPath)
      .withConverter(eventConverter)
  }

  private getEventDocRef(eventId: string) {
    return this.getEventColRef()
      .doc(eventId)
      .withConverter(eventConverter)
  }

  private getEventVoteDocRef({ eventId, currentUid }: { eventId: string, currentUid: string }) {
    const docPath = projectConfig.events.firestoreEventVoteDoc(eventId, currentUid)
    return firebase.firestore().doc(docPath).withConverter(eventVoteConverter)
  }

  private getEventSignupDoc({ eventId, currentUid }: { eventId: string, currentUid: string }) {
    const docPath = projectConfig.events.firestoreEventSignupDoc(eventId, currentUid)
    return firebase.firestore().doc(docPath).withConverter(eventSignupConverter)
  }
}

const eventConverter = {
  toFirestore(event: Event): firebase.firestore.DocumentData {
    return event;
  },
  fromFirestore(
    snapshot: firebase.firestore.QueryDocumentSnapshot,
    options: firebase.firestore.SnapshotOptions
  ): Event {
    const data: any = snapshot.data(options)!;
    return {
      ...data,
      id: snapshot.id
    }
  }
}

const eventVoteConverter = {
  toFirestore(event: EventTimeVotes): firebase.firestore.DocumentData {
    return event;
  },
  fromFirestore(
    snapshot: firebase.firestore.QueryDocumentSnapshot,
    options: firebase.firestore.SnapshotOptions
  ): EventTimeVotes {
    const data: any = snapshot.data(options)!;
    return {
      ...data,
      id: snapshot.id
    }
  }
}

const eventSignupConverter = {
  toFirestore(event: EventSignup): firebase.firestore.DocumentData {
    return event;
  },
  fromFirestore(
    snapshot: firebase.firestore.QueryDocumentSnapshot,
    options: firebase.firestore.SnapshotOptions
  ): EventSignup {
    const data: any = snapshot.data(options)!;
    return {
      ...data,
      id: snapshot.id
    }
  }
}


export const eventRepositoryFirestore = new EventRepositoryFirestore()
