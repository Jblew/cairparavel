<template>
  <create-event-panels :machine="machine" :state="state" />
</template>
<script lang="ts">
import { Component, Prop, Vue, Provide } from 'vue-property-decorator';
import { AuthenticatedLayout } from '@/features/layout';
import { CreateEventInterpreter, createEventMachine } from './machine'
import { interpret } from 'xstate';
import { CreateEventPanels } from './panels'

@Component({
  components: {
    CreateEventPanels,
  }
})
export default class CreateForm extends Vue {
  machine: CreateEventInterpreter = interpret(createEventMachine)

  state: CreateEventInterpreter['state'] = createEventMachine.initialState

  beforeMount() {
    this.machine.onTransition((state) => {
      console.log(state)
      this.state = state
    }).start()
  }

  beforeDestroy() {
    this.machine.stop()
  }
}
</script>

<style scoped>
</style>
