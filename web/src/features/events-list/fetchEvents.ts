import { eventRepository } from '@/repository'

export async function fetchEvents() {
  return eventRepository.fetchEvents()
}
