import { EventMachineContext } from '@/businesslogic'
import { eventRepositoryFirestore } from '@/repository'

export function syncEventActorFactory(eventId: string) {
  return (
    send: any,
    _onReceive: any,
  ) => {
    const unsubscribeFn = eventRepositoryFirestore.subscribeToEvent({
      eventId, onUpdated(event) { send({ type: 'UPDATED', event }) }, onError(err) {
        send({ type: 'SYNC_ERROR' })
        console.error('Sync error', err)
      }
    })
    return () => {
      unsubscribeFn()
    }
  }
}

export async function timeVote(
  { eventId, currentUid }: EventMachineContext,
  { time }: { time: number } | any,
) {
  return eventRepositoryFirestore.timeVote({ eventId, currentUid, time })
}

export async function timeUnvote(
  { eventId, currentUid }: EventMachineContext,
  { time }: { time: number } | any,
) {
  return eventRepositoryFirestore.timeUnvote({ eventId, currentUid, time })
}

export async function timeConfirm(
  { eventId }: EventMachineContext,
  { startTime, endTime }: { startTime: number; endTime: number } | any,
  evt: { startTime: number; endTime: number } | any,
) {
  return eventRepositoryFirestore.timeConfirm({ eventId, startTime, endTime })
}

export async function memberSignup({ eventId, currentUid }: EventMachineContext) {
  return eventRepositoryFirestore.memberSignup({ eventId, currentUid })
}

export async function memberSignout({ eventId, currentUid }: EventMachineContext) {
  return eventRepositoryFirestore.memberSignout({ eventId, currentUid })
}

export async function deleteEvent({ eventId }: EventMachineContext) {
  return eventRepositoryFirestore.deleteEvent(eventId)
}

export async function updateDetails(
  { eventId }: EventMachineContext,
  { name, description }: { name: string; description: string } | any,
) {
  return eventRepositoryFirestore.updateDetails({ eventId, description, name })
}
