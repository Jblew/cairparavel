<template>
  <span v-if="hasError">Error checking role {{ role }}</span>
  <span v-else-if="hasRole">Is {{ role }}</span>
  <span v-else>Is not {{ role }}</span>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { projectConfig } from '@/config';
import { checkHasRole } from './checkHasRole'

@Component
export default class extends Vue {
  @Prop()
  role: string

  @Prop()
  uid: string

  hasRole = false
  isError = false

  mounted() {
    checkHasRole(this.uid, this.role).catch(err => {
      this.isError = true
      this.hasRole = false
      console.error(err)
    })
  }
}
</script>

<style scoped>
</style>
