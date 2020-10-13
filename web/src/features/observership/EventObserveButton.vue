<template>
  <span>
    <b-button v-if="isObserving" size="sm" @click="unobserve()">
      Unobserve
    </b-button>
    <b-button v-else size="sm" @click="observe()">
      Observe
    </b-button>
    <span v-if="error">
      Error
    </span>
  </span>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { Event, EventMachineInterpreter, EventComment } from '@/businesslogic'
import { Resource } from 'vue-stateful-resource'
import { eventObserverRepositoryFirestore } from '@/repository'

@Component
export default class extends Vue {
  @Prop({ required: true })
  event!: Event

  observingRes: Resource<{ observing: boolean }> = Resource.success({ observing: false })
  unsubscribeFn?: () => void

  error = false

  beforeMount() {
    const eventId = this.event.id
    if (!eventId) throw new Error('Missing event.id')
    this.unsubscribeFn?.()
    this.unsubscribeFn = eventObserverRepositoryFirestore.subscribeToMyObservership({
      eventId,
      on: (res) => { this.observingRes = res }
    })
  }

  beforeDestroy() {
    this.unsubscribeFn?.()
  }

  get isObserving() {
    return this.observingRes.result && this.observingRes.result.observing
  }

  async observe() {
    try {
      this.error = false
      const eventId = this.event.id
      if (!eventId) throw new Error('Missing event.id')
      await eventObserverRepositoryFirestore.observe(eventId)
    }
    catch (err) {
      console.error(err)
      this.error = true
    }
  }

  async unobserve() {
    try {
      this.error = false
      const eventId = this.event.id
      if (!eventId) throw new Error('Missing event.id')
      await eventObserverRepositoryFirestore.unobserve(eventId)
    }
    catch (err) {
      console.error(err)
      this.error = true
    }
  }
}
</script>

<style scoped></style>
