<template>
  <div data-test="panel-participant-limits">
    <h2>Set participant limits</h2>

    <p>
      Min participants:
      <input
        type="number"
        v-model="minParticipants"
        data-test="input-min-participants"
      />
    </p>

    <p>
      Max participants:
      <input
        type="number"
        v-model="maxParticipants"
        data-test="input-max-participants"
      />
    </p>

    <p><button @click="next()" data-test="btn-next">Next</button></p>

    <p><button @click="back()" data-test="btn-back">Back</button></p>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { CreateEventInterpreter } from '../machine'

@Component
export default class extends Vue {
  @Prop()
  machine: CreateEventInterpreter

  @Prop()
  state: CreateEventInterpreter['state']


  minParticipants = 0
  maxParticipants = 100

  mounted() {
    this.minParticipants = this.state.context.event.minParticipants
    this.maxParticipants = this.state.context.event.maxParticipants
  }

  next() {
    this.machine.send({
      type: 'SET_PARTICIPANT_LIMITS',
      limits: { minParticipants: Number(this.minParticipants), maxParticipants: Number(this.maxParticipants) }
    })
  }

  back() {
    this.machine.send('BACK')
  }
}
</script>

<style scoped></style>
