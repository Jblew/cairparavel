import { Event } from '@/businesslogic';

import { eventRepositoryFirestore } from '@/repository';

export function saveEvent(event: Event) {
  return eventRepositoryFirestore.saveEvent(event)
}
