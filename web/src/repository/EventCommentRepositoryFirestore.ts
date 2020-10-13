import { projectConfig } from '@/config'
import firebase from 'firebase/app'
import { Event, EventComment, EventCommentRepository, EventSignup, EventTimeVotes, validateEvent } from '@/businesslogic'
import { Resource } from 'vue-stateful-resource'
import { listenForFirestoreSnapshots } from '@/util'

export class EventCommentRepositoryFirestore implements EventCommentRepository {
  subscribe({ eventId, on }: { eventId: string, on: (res: Resource<EventComment[]>) => void }) {
    const query = firebase.firestore()
      .collection(projectConfig.comments.firestoreEventCommentsCol(eventId))
      .withConverter(eventComment)
      .orderBy('time', 'desc')
    const unsubscribeFn = listenForFirestoreSnapshots<EventComment>(
      query,
      (snapshots) => snapshots.map((snapshot) => snapshot.data() as EventComment),
      on,
    )
    return unsubscribeFn
  }

  add(eventId: string, contents: string): Promise<void> {
    const user = firebase.auth().currentUser!
    const comment: EventComment = {
      eventId,
      authorUid: user.uid,
      authorDisplayName: user.displayName || 'No name',
      contents,
      time: Date.now(),
    }
    return firebase.firestore()
      .collection(projectConfig.comments.firestoreEventCommentsCol(comment.eventId))
      .withConverter(eventComment)
      .doc()
      .set(comment)
  }
}

const eventComment = {
  toFirestore(comment: EventComment): firebase.firestore.DocumentData {
    return {
      ...comment,
      time: Date.now(),
    }
  },
  fromFirestore(
    snapshot: firebase.firestore.QueryDocumentSnapshot,
    options: firebase.firestore.SnapshotOptions
  ): EventComment {
    const data: any = snapshot.data(options)!;
    return {
      ...data,
      id: snapshot.id,
    }
  }
}

export const eventCommentRepositoryFirestore = new EventCommentRepositoryFirestore()
