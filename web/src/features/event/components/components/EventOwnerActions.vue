<template>
  <span class="event-owner-actions">
    <event-details-edit-modal
      v-if="canEdit"
      :event="event"
      @save="saveDetails"
    />
    &nbsp;
    <b-button
      v-if="canDelete"
      v-b-modal="'modal-delete-event-' + event.id"
      size="sm"
    >
      Delete
    </b-button>
    <b-modal
      :id="'modal-delete-event-' + event.id"
      :title="'Delete event ' + event.name"
      header-bg-variant="danger"
      ok-variant="danger"
      @ok="deleteEvent()"
    >
      Delete event
    </b-modal>
  </span>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { Event, EventMachineInterpreter } from '@/businesslogic'
import EventLink from './EventLink.vue'
import EventOwnerActions from './EventOwnerActions.vue'
import EventDetailsEditModal from './EventDetailsEditModal.vue'

@Component({
  components: {
    EventLink,
    EventOwnerActions,
    EventDetailsEditModal,
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

  get canDelete(): boolean {
    return this.state.nextEvents.includes('DELETE')
  }

  get canEdit(): boolean {
    return this.state.nextEvents.includes('UPDATE_DETAILS')
  }

  deleteEvent() {
    this.interpreter.send('DELETE')
  }

  saveDetails(details: { description: string, name: string }) {
    this.interpreter.send({ type: 'UPDATE_DETAILS', ...details })
  }
}
</script>

<style scoped>
.event-owner-actions {
  font-size: 75%;
  padding-left: 0.5rem;
}
</style>
