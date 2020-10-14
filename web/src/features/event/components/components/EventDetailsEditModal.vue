<template>
  <span>
    <b-button v-b-modal="'modal-edit-event-' + event.id" size="sm">
      Edit
    </b-button>
    <b-modal
      :id="'modal-edit-event-' + event.id"
      :title="'Edit event ' + event.name"
      @show="resetModel()"
      @ok="saveEvent()"
    >
      <event-details-edit-form v-model="model" />
    </b-modal>
  </span>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { Event, EventMachineInterpreter } from '@/businesslogic'
import EventLink from './EventLink.vue'
import EventDetailsEditForm from './EventDetailsEditForm.vue'

@Component({
  components: {
    EventDetailsEditForm,
  }
})
export default class extends Vue {
  @Prop({ required: true })
  event: Event

  model = { name: "", description: "" }

  resetModel() {
    this.model.name = this.event.name
    this.model.description = this.event.description
  }

  saveEvent() {
    this.$emit('save', this.model)
  }
}
</script>

<style scoped>
.event-owner-actions {
  font-size: 75%;
  padding-left: 0.5rem;
}
</style>
