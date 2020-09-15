<template>
  <span>
    <span v-for="stateName in stateNames" :key="stateName">
      <slot :name="stateName"></slot>
    </span>
  </span>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { State } from 'xstate'

@Component
export default class extends Vue {
  @Prop({ required: true })
  state!: State<any, any>

  get stateNames() {
    console.log(this.state.toStrings())
    return this.state.toStrings()
  }
}
</script>
