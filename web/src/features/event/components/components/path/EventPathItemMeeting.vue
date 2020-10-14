<template>
  <event-path-item :enabled="enabled" :checked="checked" name="Meeting">
    Starts {{ event.startTime | formatTimeMillis }}, Duration
    {{ duration | formatInvervalMsToH }}
    <span v-if="inProggress">In proggress</span>
  </event-path-item>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { Event } from '@/businesslogic'
import EventPathItem from './EventPathItem.vue'

function formatTimeMillis(timeMs: number): string {
  return new Date(timeMs).toISOString()
}

function formatInvervalMsToH(intervalMs: number): string {
  const hours = intervalMs / 1000 / 3600
  return hours.toFixed(1) + 'h'
}

@Component({
  components: {
    EventPathItem,
  },
  filters: {
    formatTimeMillis,
    formatInvervalMsToH,
  }
})
export default class extends Vue {
  @Prop({ required: true })
  event: Event

  @Prop({ required: true })
  enabled: boolean

  @Prop({ required: true })
  checked: boolean

  @Prop({ default: () => false })
  inProggress: boolean

  get duration(): number {
    return this.event.endTime - this.event.startTime
  }
}
</script>

<style scoped></style>
