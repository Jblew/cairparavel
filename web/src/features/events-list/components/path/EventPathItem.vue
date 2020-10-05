<template>
  <div :class="{ pathitem: true, enabled, disabled: !enabled }">
    <div class="level">
      <span v-if="checked">[x]</span>
      <span v-else>[ ]</span>
      {{ name }}
    </div>
    <div class="level">
      <slot />
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'

@Component
export default class extends Vue {
  @Prop({ required: true })
  name: string

  @Prop({ required: true })
  checked: boolean

  @Prop({ required: true })
  enabled: boolean
}
</script>

<style scoped>
.pathitem {
  display: inline-block;
  border: 1px solid #ddd;
  margin: 0.5rem;
}

.pathitem .level {
  width: 100%;
  display: block;
}

.disabled {
  color: gray;
}
</style>
