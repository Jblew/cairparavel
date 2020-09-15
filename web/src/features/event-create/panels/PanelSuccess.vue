<template>
  <div data-test="panel-success">
    <h2>Successfully created event</h2>
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
