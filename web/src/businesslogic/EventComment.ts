import { Resource } from 'vue-stateful-resource';

export interface EventComment {
  id?: string
  eventId: string
  authorUid: string
  authorDisplayName: string
  contents: string
  time: number
}

export interface EventCommentRepository {
  subscribe(opts: { eventId: string, on: (res: Resource<EventComment[]>) => void }): () => void
  add(eventId: string, comment: string): Promise<void>
}
