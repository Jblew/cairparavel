<template>
  <div data-test="panel-details">
    <h2>Details setup</h2>
    <input type="text" data-test="input-name" />
    <textarea data-test="input-description"></textarea>
    <button data-test="btn-next">Next</button>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import { CreateEventInterpreter } from '../machine'

@Component
export default class extends Vue {
  @Inject()
  machine: CreateEventInterpreter

  @Inject()
  state: CreateEventInterpreter['state']
}
</script>

<style scoped>
</style>
