<template>
  <div data-test="panel-participant-limits">
    <h2>Set participant limits</h2>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import { CreateEventInterpreter } from '../machine'

@Component
export default class extends Vue {
  @Prop()
  machine: CreateEventInterpreter

  @Prop()
  state: CreateEventInterpreter['state']
}
</script>

<style scoped>
</style>
