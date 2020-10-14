<template>
  <event-path-item :enabled="enabled" :checked="checked" name="Meeting">
    Time will be confirmed
  </event-path-item>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { Event } from '@/businesslogic'
import EventPathItem from './EventPathItem.vue'

function formatTimeMillis(date: Date): string {
  return date.toISOString()
}

function formatInvervalMsToH(intervalMs: number): string {
  const hours = intervalMs / 1000 / 3600
  return hours.toFixed(1) + 'h'
}

@Component({
  components: {
    EventPathItem,
  }
})
export default class extends Vue {
  @Prop({ required: true })
  event: Event

  @Prop({ required: true })
  enabled: boolean

  @Prop({ required: true })
  checked: boolean
}
</script>

<style scoped></style>
