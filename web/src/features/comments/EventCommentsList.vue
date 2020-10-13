<template>
  <div class="event-comments">
    <stateful-resource :resource="commentsResource">
      <ul>
        <li v-for="comment in comments" :key="comment.id">
          <event-comment-item :comment="comment" />
        </li>
      </ul>
    </stateful-resource>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { Event, EventMachineInterpreter, EventComment } from '@/businesslogic'
import { StatefulResource, Resource } from 'vue-stateful-resource'
import EventCommentItem from './EventCommentItem.vue'
import { eventCommentRepositoryFirestore } from '@/repository'

@Component({
  components: {
    StatefulResource,
    EventCommentItem,
  }
})
export default class extends Vue {
  @Prop({ required: true })
  event!: Event

  commentsResource: Resource<EventComment[]> = Resource.empty()

  unsubscribeFn?: () => void

  beforeMount() {
    const eventId = this.event.id
    if (!eventId) throw new Error('Missing event.id')
    this.unsubscribeFn?.()
    this.unsubscribeFn = eventCommentRepositoryFirestore.subscribe({
      eventId,
      on: (res) => this.commentsResource = res
    })
  }

  beforeDestroy() {
    this.unsubscribeFn?.()
  }

  get comments(): EventComment[] {
    return this.commentsResource.result || [] as EventComment[]
  }
}
</script>

<style scoped></style>
