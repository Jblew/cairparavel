<template>
  <b-form>
    <b-form-group
      label="Name:"
      label-for="input-name"
      description="Choose descriptive event name"
    >
      <b-form-input
        id="input-name"
        v-model="name"
        type="text"
        required
        placeholder="Enter email"
      ></b-form-input>
    </b-form-group>

    <b-form-textarea
      id="textarea"
      v-model="description"
      placeholder="Description of the event"
      rows="7"
      max-rows="15"
    ></b-form-textarea>
  </b-form>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'

@Component
export default class extends Vue {
  @Prop()
  value: {
    name: string
    description: string
  }

  get name(): string {
    return this.value.name
  }

  set name(name: string) {
    this.$emit('input', { ...this.value, name })
  }

  get description(): string {
    return this.value.description
  }

  set description(description: string) {
    this.$emit('input', { ...this.value, description })
  }
}

</script>

<style scoped>
</style>
