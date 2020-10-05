<template>
  <span>
    <event-header :event="event" />
    <event-path>
      <event-path-item-created :event="event" />
      <event-path-separator />
      <event-path-item :enabled="true" :checked="true" name="Signup">
        TODO List of members
      </event-path-item>
      <event-path-separator />
      <event-path-item :enabled="true" :checked="true" name="Cancelled" />
      <event-path-separator />
      <event-path-item :enabled="false" :checked="false" name="Meeting">
        TODO meeting time
      </event-path-item>
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
} from '../components'

@Component({
  components: {
    EventHeader,
    EventPath,
    EventPathItem,
    EventPathSeparator,
    EventPathItemCreated,
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
