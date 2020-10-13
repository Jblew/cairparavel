<template>
  <event-path-item :enabled="true" :checked="true" name="Created">
    by {{ ownerDisplayName }}
  </event-path-item>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { Event } from '@/businesslogic'
import EventPathItem from './EventPathItem.vue'

@Component({
  components: {
    EventPathItem,
  }
})
export default class extends Vue {
  @Prop({ required: true })
  event: Event

  get ownerDisplayName(): string {
    return this.event.ownerDisplayName
  }
}
</script>

<style scoped></style>
