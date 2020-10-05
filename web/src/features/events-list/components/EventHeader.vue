<template>
  <event-link :event="event">
    <strong>{{ event.name }}</strong>
  </event-link>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { Event } from '@/businesslogic'
import EventLink from './EventLink.vue'

@Component({
  components: {
    EventLink,
  }
})
export default class extends Vue {
  @Prop({ required: true })
  event!: Event
}
</script>

<style scoped></style>
