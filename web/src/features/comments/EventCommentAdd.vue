<template>
  <div class="event-comment-add">
    <textarea v-model="contents" data-test="input-contents"></textarea>
    <button data-test="btn-add" @click="add()">
      Comment
    </button>
    <span v-if="loading">Commenting...</span>
    <span v-if="error">Error</span>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { Event, EventMachineInterpreter, EventComment } from '@/businesslogic'
import { StatefulResource, Resource } from 'vue-stateful-resource'
import { eventCommentRepositoryFirestore } from '@/repository'

@Component
export default class extends Vue {
  @Prop({ required: true })
  event!: Event

  contents = ""

  loading = false
  error = false

  async add() {
    this.loading = true
    this.error = false
    try {
      const eventId = this.event.id
      if (!eventId) throw new Error('Missing event.id')
      await eventCommentRepositoryFirestore.add(eventId, this.contents)
    }
    catch (error) {
      console.error(error)
      this.error = true
    }
    this.loading = false
    this.contents = ""
  }

}
</script>

<style scoped></style>
