<template>
  <div data-test="panel-voting-question">
    <h2>Do you want members to vote on event time?</h2>

    <p>
      <input
        type="checkbox"
        :checked="isVoting"
        @click="selectVoting()"
        data-test="btn-choose-voting"
      />
      Allow members vote
    </p>
    <p>
      <input
        type="checkbox"
        :checked="isNoVoting"
        @click="selectNoVoting()"
        data-test="btn-choose-no-voting"
      />
      Set arbitrary time
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

  get isVoting() {
    return this.state.matches('QuestionVoting.Voting')
  }

  get isNoVoting() {
    return this.state.matches('QuestionVoting.NoVoting')
  }

  selectVoting() {
    this.machine.send('CHOOSE_VOTING')
  }

  selectNoVoting() {
    this.machine.send('CHOOSE_NO_VOTING')
  }

  next() {
    this.machine.send('NEXT')
  }

  back() {
    this.machine.send('BACK')
  }
}
</script>

<style scoped></style>
