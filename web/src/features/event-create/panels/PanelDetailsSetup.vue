<template>
  <div data-test="panel-details">
    <h2>Details setup</h2>
    <input type="text" v-model="name" data-test="input-name" />
    <textarea v-model="description" data-test="input-description"></textarea>
    <button data-test="btn-next" @click="save()">Next</button>
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

  name = ''
  description = ''

  mounted() {
    this.name = this.state.context.event.name
    this.description = this.state.context.event.description
  }

  save() {
    this.machine.send({
      type: 'SET_DETAILS',
      details: { name: this.name, description: this.description },
    })
  }
}
</script>

<style scoped></style>
