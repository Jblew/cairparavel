<template>
  <span class="event-owner-actions">
    <b-button
      v-if="canEdit"
      v-b-modal="'modal-edit-event-' + event.id"
      size="sm"
    >
      Edit
    </b-button>
    &nbsp;
    <b-button
      v-if="canDelete"
      v-b-modal="'modal-delete-event-' + event.id"
      size="sm"
    >
      Delete
    </b-button>

    <b-modal
      :id="'modal-edit-event-' + event.id"
      :title="'Edit event ' + event.name"
    >
      Edit event
    </b-modal>

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

@Component({
  components: {
    EventLink,
    EventOwnerActions,
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
    console.log('NEXT_EVENTS: ', this.state.nextEvents)
    return this.state.nextEvents.includes('DELETE')
  }

  get canEdit(): boolean {
    console.log('NEXT_EVENTS: ', this.state.nextEvents)
    return this.state.nextEvents.includes('UPDATE_DETAILS')
  }

  deleteEvent() {
    this.interpreter.send('DELETE')
  }
}
</script>

<style scoped>
.event-owner-actions {
  font-size: 75%;
  padding-left: 0.5rem;
}
</style>
