<template>
  <div class="event-comments">
    <event-comments-list :event="event" />
    <event-comment-add :event="event" />
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { Event, EventMachineInterpreter, EventComment } from '@/businesslogic'
import EventCommentAdd from './EventCommentAdd.vue'
import EventCommentsList from './EventCommentsList.vue'

@Component({
  components: {
    EventCommentAdd,
    EventCommentsList
  }
})
export default class extends Vue {
  @Prop({ required: true })
  event!: Event
}
</script>

<style scoped></style>
