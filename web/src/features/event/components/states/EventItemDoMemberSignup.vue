<template>
  <span>
    <event-header :interpreter="interpreter" :state="state" :event="event" />
    signing in...
  </span>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { Event, EventMachineInterpreter } from '@/businesslogic'
import { EventHeader } from '../components'

@Component({
  components: {
    EventHeader,
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
