import { projectConfig } from '@/config'
import firebase from 'firebase/app'
import { EventObserver, EventObserverRepository } from '@/businesslogic'
import { Resource } from 'vue-stateful-resource'

export class EventObserverRepositoryFirestore implements EventObserverRepository {
  subscribeToMyObservership({ eventId, on }: { eventId: string, on: (res: Resource<{ observing: boolean }>) => void }): () => void {
    on(Resource.loading());

    const uid = firebase.auth().currentUser!.uid
    const query = firebase.firestore().doc(projectConfig.events.observerDoc(eventId, uid))
    const unsubscribeFn = query.onSnapshot(
      (snapshot) => {
        on(Resource.success({ observing: snapshot.exists }));
      },
      (error) => {
        on(Resource.error(error.message));
      },
    );
    return unsubscribeFn
  }

  async observe(eventId: string): Promise<void> {
    const uid = firebase.auth().currentUser!.uid
    const path = projectConfig.events.observerDoc(eventId, uid)
    const observer: EventObserver = { uid, eventId }
    await firebase.firestore().doc(path).set(observer)
  }

  async unobserve(eventId: string): Promise<void> {
    const uid = firebase.auth().currentUser!.uid
    await firebase.firestore().doc(projectConfig.events.observerDoc(eventId, uid)).delete()
  }
}


export const eventObserverRepositoryFirestore = new EventObserverRepositoryFirestore()
