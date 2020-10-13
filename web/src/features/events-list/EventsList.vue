<template>
  <stateful-resource :resource="eventsResource">
    <ul data-test="events-list">
      <li v-for="event in events" :key="event.id">
        <event-item :event="event" />
      </li>
    </ul>
  </stateful-resource>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { StatefulResource, Resource } from 'vue-stateful-resource'
import { fetchEvents } from './fetchEvents'
import { Event } from '@/businesslogic'
import { EventItem } from '@/features/event'

@Component({
  components: {
    StatefulResource,
    EventItem,
  },
})
export default class EventsList extends Vue {
  eventsResource: Resource<Event[]> = Resource.empty()

  mounted() {
    Resource.fetchResource(
      'events',
      fetchEvents,
      res => (this.eventsResource = res),
    )
  }

  get events(): Event[] {
    return this.eventsResource.result || []
  }
}
</script>

<style scoped>
ul li {
  text-align: left;
}
</style>
