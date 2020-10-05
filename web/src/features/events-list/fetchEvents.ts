import { eventRepositoryFirestore } from '@/repository'

export async function fetchEvents() {
  return eventRepositoryFirestore.fetchEvents()
}
