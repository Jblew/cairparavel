import { EventMachineContext } from '@/businesslogic'
import { eventRepository } from '@/repository'

export function syncEventActorFactory(eventId: string) {
  return (_ctx: EventMachineContext, _evt: any) => (
    send: any,
    _onReceive: any,
  ) => {
    const unsubscribeFn = eventRepository.subscribeToEvent({
      eventId, onUpdated(event) { send({ type: 'UPDATED', event }) }, onError(err) {
        send({ type: 'SYNC_ERROR' })
        console.error('Sync error', err)
      }
    })
    return () => {
      console.log('UNSUBSCRIBE')
      unsubscribeFn()
    }
  }
}

export async function timeVote(
  { eventId, currentUid }: EventMachineContext,
  { time }: { time: number } | any,
) {
  return eventRepository.timeVote({ eventId, currentUid, time })
}

export async function timeUnvote(
  { eventId, currentUid }: EventMachineContext,
  { time }: { time: number } | any,
) {
  return eventRepository.timeUnvote({ eventId, currentUid, time })
}

export async function timeConfirm(
  { eventId }: EventMachineContext,
  { startTime, endTime }: { startTime: number; endTime: number } | any,
  evt: { startTime: number; endTime: number } | any,
) {
  return eventRepository.timeConfirm({ eventId, startTime, endTime })
}

export async function memberSignup({ eventId, currentUid }: EventMachineContext) {
  return eventRepository.memberSignup({ eventId, currentUid })
}

export async function memberSignout({ eventId, currentUid }: EventMachineContext) {
  return eventRepository.memberSignout({ eventId, currentUid })
}

export async function updateDetails(
  { eventId }: EventMachineContext,
  { name, description }: { name: string; description: string } | any,
) {
  return eventRepository.updateDetails({ eventId, description, name })
}
