import { projectConfig } from '@/config'
import firebase from 'firebase/app'
import { EventMachineContext, Event, EventTimeVotes } from '@/businesslogic'

export function syncEventActorFactory(eventId: string) {
  return (_ctx: EventMachineContext, _evt: any) => (
    send: any,
    _onReceive: any,
  ) => {
    const docPath = projectConfig.events.firestoreEventDoc(eventId)
    const unsubscribeFn = firebase
      .firestore()
      .doc(docPath)
      .onSnapshot(
        doc => send({ type: 'UPDATED', event: doc.data() }),
        err => {
          send({ type: 'SYNC_ERROR' })
          console.error('Sync error', err)
        },
      )
    return unsubscribeFn
  }
}

export async function timeVote(
  ctx: EventMachineContext,
  evt: { time: number } | any,
) {
  const docPath = projectConfig.events.firestoreEventVoteDoc(
    ctx.eventId,
    ctx.currentUid,
  )
  const docRef = firebase.firestore().doc(docPath)
  const displayName = firebase.auth().currentUser!.displayName || 'Unknown'
  const currentVoteSnapshot = await docRef.get()
  const times: number[] = currentVoteSnapshot.exists
    ? [...currentVoteSnapshot.data()!.times, evt.time]
    : [evt.time]
  const votes: EventTimeVotes = {
    displayName,
    times,
  }
  return docRef.set(votes)
}

export async function timeUnvote(
  ctx: EventMachineContext,
  evt: { time: number } | any,
) {
  const docPath = projectConfig.events.firestoreEventVoteDoc(
    ctx.eventId,
    ctx.currentUid,
  )
  const docRef = firebase.firestore().doc(docPath)
  const displayName = firebase.auth().currentUser!.displayName || 'Unknown'
  const currentVoteSnapshot = await docRef.get()
  const times: number[] = currentVoteSnapshot.exists
    ? [
      ...currentVoteSnapshot
        .data()!
        .times.filter((time: any) => time != evt.time),
    ]
    : []
  const votes: EventTimeVotes = {
    displayName,
    times,
  }
  return docRef.set(votes)
}

export async function timeConfirm(
  ctx: EventMachineContext,
  evt: { startTime: number; endTime: number } | any,
) {
  const docPath = projectConfig.events.firestoreEventDoc(ctx.eventId)
  const docRef = firebase.firestore().doc(docPath)
  const update: Partial<Event> = {
    timeConfirmed: true,
    startTime: evt.startTime,
    endTime: evt.endTime,
  }
  return docRef.set(update)
}

export async function memberSignup(ctx: EventMachineContext) {
  const docPath = projectConfig.events.firestoreEventSignupDoc(
    ctx.eventId,
    ctx.currentUid,
  )
  const docRef = firebase.firestore().doc(docPath)
  const displayName = firebase.auth().currentUser!.displayName || 'Unknown'
  return docRef.set({ displayName })
}

export async function memberSignout(ctx: EventMachineContext) {
  const docPath = projectConfig.events.firestoreEventSignupDoc(
    ctx.eventId,
    ctx.currentUid,
  )
  const docRef = firebase.firestore().doc(docPath)
  return docRef.delete()
}

export async function updateDetails(
  ctx: EventMachineContext,
  evt: { name: string; description: string } | any,
) {
  const docPath = projectConfig.events.firestoreEventSignupDoc(
    ctx.eventId,
    ctx.currentUid,
  )
  const docRef = firebase.firestore().doc(docPath)
  return docRef.set({
    name: evt.name,
    description: evt.description,
  })
}
