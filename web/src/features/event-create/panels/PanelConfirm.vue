<template>
  <div data-test="panel-confirm">
    <h2>Confirm</h2>

    <p><button @click="save()" data-test="btn-next">Save</button></p>

    <p><button @click="back()" data-test="btn-back">Back</button></p>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { CreateEventInterpreter } from '../machine'

@Component
export default class extends Vue {
  @Prop()
  machine: CreateEventInterpreter

  @Prop()
  state: CreateEventInterpreter['state']

  save() {
    this.machine.send('DO_SAVE')
  }

  back() {
    this.machine.send('BACK')
  }
}
</script>

<style scoped></style>
