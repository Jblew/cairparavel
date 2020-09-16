<template>
  <div data-test="panel-event-time-setup">
    <h2>Setup event time</h2>

    <p>
      Event start time:
      <date-time-input v-model="startTime" data-test="input-start-time" />
    </p>
    <p>
      Event end time:
      <date-time-input v-model="endTime" data-test="input-end-time" />
    </p>
    <p><button @click="next()" data-test="btn-next">Next</button></p>

    <p><button @click="back()" data-test="btn-back">Back</button></p>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { CreateEventInterpreter } from '../machine'
import { DateTimeInput } from '@/components'

@Component({
  components: {
    DateTimeInput,
  }
})
export default class extends Vue {
  @Prop()
  machine: CreateEventInterpreter

  @Prop()
  state: CreateEventInterpreter['state']

  startTime = Date.now()
  endTime = Date.now()

  mounted() {
    this.startTime = this.state.context.event.startTime
    this.endTime = this.state.context.event.endTime
  }

  next() {
    this.machine.send({
      type: 'SET_EVENT_TIME',
      time: { startTime: this.startTime, endTime: this.endTime }
    })
  }

  back() {
    this.machine.send('BACK')
  }
}
</script>

<style scoped></style>
