<template>
  <span>
    <event-header :event="event" />
    Waiting for time confirmation
  </span>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { Event, EventMachineInterpreter } from '@/businesslogic'

@Component({
  components: {}
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
