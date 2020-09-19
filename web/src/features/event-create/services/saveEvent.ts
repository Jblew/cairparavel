import { Event } from '@/businesslogic';

import { eventRepository } from '@/repository';

export function saveEvent(event: Event) {
  return eventRepository.saveEvent(event)
}
