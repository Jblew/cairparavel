<template>
  <router-link :to="eventUrl">
    <slot />
  </router-link>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { Event } from '@/businesslogic'
import { projectConfig } from '@/config'

@Component
export default class extends Vue {
  @Prop({ required: true })
  event!: Event

  get eventUrl() {
    return projectConfig.events.eventUrl(this.event)
  }
}
</script>

<style scoped></style>
