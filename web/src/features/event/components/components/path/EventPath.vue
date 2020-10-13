<template>
  <div class="eventpath">
    <slot />
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'

@Component
export default class extends Vue {

}

</script>

<style scoped>
.eventpath {
  display: block;
  background: #eee;
}
</style>
