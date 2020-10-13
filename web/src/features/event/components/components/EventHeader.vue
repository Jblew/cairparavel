<template>
  <span>
    <strong>{{ event.name }}</strong>
    <event-owner-actions
      v-if="isEventOwner"
      :interpreter="interpreter"
      :state="state"
    />
  </span>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { Event, EventMachineInterpreter } from '@/businesslogic'
import EventLink from './EventLink.vue'
import EventOwnerActions from './EventOwnerActions.vue'

@Component({
  components: {
    EventLink,
    EventOwnerActions,
  }
})
export default class extends Vue {
  @Prop({ required: true })
  event!: Event

  @Prop({ required: true })
  interpreter: EventMachineInterpreter

  @Prop({ required: true })
  state: EventMachineInterpreter['state']

  get isEventOwner(): boolean {
    return this.state.context.currentUid === this.state.context.event?.ownerUid
  }
}
</script>

<style scoped></style>
