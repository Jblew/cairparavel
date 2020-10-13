import { Resource } from 'vue-stateful-resource';

export interface EventObserver {
  eventId: string
  uid: string
}

export interface EventObserverRepository {
  subscribeToMyObservership(opts: { eventId: string, on: (res: Resource<{ observing: boolean }>) => void }): () => void
  observe(eventId: string): Promise<void>
  unobserve(eventId: string): Promise<void>
}
