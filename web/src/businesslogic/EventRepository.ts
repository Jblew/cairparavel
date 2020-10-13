import { Event } from './Event'

export interface EventRepository {
  subscribeToEvent(optd: { eventId: string, onUpdated: (e: Event) => void, onError: (err: Error) => void }): () => void
  saveEvent(event: Event): Promise<string>
  fetchEvents(): Promise<Event[]>
  timeVote(opts: { eventId: string, currentUid: string, time: number }): Promise<void>
  timeUnvote(opts: { eventId: string, currentUid: string, time: number }): Promise<void>
  timeConfirm(
    opts: { eventId: string, startTime: number; endTime: number }
  ): Promise<void>
  memberSignup({ eventId, currentUid }: { eventId: string, currentUid: string }): Promise<void>
  memberSignout({ eventId, currentUid }: { eventId: string, currentUid: string }): Promise<void>
  updateDetails(
    { eventId, name, description }: { eventId: string, name: string; description: string },
  ): Promise<void>
  deleteEvent(eventId: string): Promise<void>
}
