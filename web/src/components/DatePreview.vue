<template>
  <div class="date-preview">
    <span>{{ start || formatDate }}</span>

    <span v-if="stop">{{ intervalMs || formatInvervalMs }}</span>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'

function formatDate(date: Date): string {
  return date.toISOString()
}

function formatInvervalMs(intervalMs: number): string {
  const hours = intervalMs / 1000 / 3600
  return hours.toFixed(1) + 'h'
}

@Component({
  filters: {
    formatDate,
    formatInvervalMs,
  }
})
export default class extends Vue {
  @Prop({ required: true })
  start: Date

  @Prop({ required: false })
  stop: Date

  get intervalMs(): number {
    return (this.stop?.getTime() || 0) - this.start.getTime()
  }
}
</script>
<style>
</style>
