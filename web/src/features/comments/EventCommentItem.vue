<template>
  <div class="event-comment">
    <strong>{{ comment.authorDisplayName }}</strong>
    &nbsp;
    <i>{{ comment.time | formatTime }}</i>
    &nbsp;
    {{ comment.contents }}
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { EventComment } from '@/businesslogic'

function formatTime(timestamp: number): string {
  return new Date(timestamp).toISOString()
}

@Component({
  filters: {
    formatTime,
  },
})
export default class extends Vue {
  @Prop({ required: true })
  comment!: EventComment
}
</script>

<style scoped></style>
