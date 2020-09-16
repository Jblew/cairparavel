<template>
  <div data-test="panel-do-save">
    <h2>Saving...</h2>
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
}
</script>

<style scoped></style>
