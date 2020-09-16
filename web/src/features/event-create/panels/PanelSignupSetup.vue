<template>
  <div data-test="panel-signup-setup">
    <h2>Signup date setup</h2>

    <p>
      Signup until:
      <date-time-input v-model="signupTime" data-test="input-signup-time" />
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

  signupTime = Date.now()

  mounted() {
    this.signupTime = this.state.context.event.signupTime
  }

  next() {
    this.machine.send({
      type: 'SET_SIGNUP_TIME',
      time: this.signupTime
    })
  }

  back() {
    this.machine.send('BACK')
  }
}
</script>

<style scoped></style>
