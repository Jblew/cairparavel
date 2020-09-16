<template>
  <div data-test="panel-voting-setup">
    <h2>Setup voting</h2>
    <p>
      <input
        type="checkbox"
        :checked="canSuggestTime"
        @click="canSuggestTime = !canSuggestTime"
        data-test="input-can-suggest-time"
      />
      Voters can suggest time
    </p>
    <p v-for="(time, index) of allowedTimes" :key="index">
      <date-time-input
        v-model="time.value"
        :data-test="'input-allowed-time-' + index"
      />
      <button @click="removeTime(index)">x</button>
    </p>
    <p>
      <button @click="addAllowedTime()" data-test="btn-add-allowed-time">
        +
      </button>
    </p>
    <p>
      Voting until:
      <date-time-input v-model="votingTime" data-test="input-voting-time" />
    </p>
    <p><button @click="next()" data-test="btn-next">Next</button></p>

    <p><button @click="back()" data-test="btn-back">Back</button></p>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
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

  canSuggestTime = false
  allowedTimes = [{ value: Date.now() }]
  votingTime = Date.now()

  mounted() {
    this.canSuggestTime = this.state.context.event.canSuggestTime
    this.allowedTimes = [
      ...this.state.context.event.allowedTimes,
      0,
    ].map(value => ({ value }))
    this.votingTime = this.state.context.event.votingTime
  }

  removeTime(index: number) {
    this.allowedTimes = this.allowedTimes.filter((_, i) => i !== index)
  }

  addAllowedTime() {
    this.allowedTimes = [...this.allowedTimes, { value: Date.now() }]
  }

  next() {
    this.machine.send({
      type: 'SET_VOTING',
      voting: {
        canSuggestTime: this.canSuggestTime,
        allowedTimes: this.allowedTimes.map(v => v.value).filter(allowed => allowed > 0),
        votingTime: this.votingTime
      }
    })
  }

  back() {
    this.machine.send('BACK')
  }
}

</script>

<style scoped></style>
