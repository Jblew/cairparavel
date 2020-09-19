<template>
  <span class="event">
    <strong class="name">
      <router-link :to="eventUrl">
        {{ event.name }}
      </router-link>
    </strong>
  </span>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
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
