<template>
  <span>
    <event-header :interpreter="interpreter" :state="state" :event="event" />
    <event-path>
      <event-path-item-created :event="event" />
      <event-path-separator />
      <event-path-item
        :enabled="true"
        :checked="true"
        name="Voting for time (closed)"
      >
        TODO voting summary
      </event-path-item>
      <event-path-separator />
      <event-path-item
        :enabled="true"
        :checked="true"
        name="Waiting for confirmation"
      >
        TODO confirmation button for owner
      </event-path-item>
      <event-path-separator />
      <event-path-item-meeting-unknown-time
        :enabled="false"
        :checked="false"
        :event="event"
      />
    </event-path>
  </span>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { Event, EventMachineInterpreter } from '@/businesslogic'
import {
  EventHeader,
  EventPath,
  EventPathItem,
  EventPathSeparator,
  EventPathItemCreated,
  EventPathItemMeetingUnknownTime,
} from '../components'

@Component({
  components: {
    EventHeader,
    EventPath,
    EventPathItem,
    EventPathItemCreated,
    EventPathSeparator,
    EventPathItemMeetingUnknownTime,
  }
})
export default class extends Vue {
  @Prop({ required: true })
  interpreter: EventMachineInterpreter

  @Prop({ required: true })
  state: EventMachineInterpreter['state']

  get event(): Event {
    return this.state.context.event!
  }
}
</script>

<style scoped></style>
