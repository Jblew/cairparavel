<template>
  <a v-if="isLoading">Revoking {{ role }}...</a>
  <a v-else-if="isError" @click="revokeRole" class="clickable"
    >Revoke {{ role }} (error!)</a
  >
  <a v-else @click="revokeRole" class="clickable">Revoke {{ role }}</a>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { revokeRole } from './helpers'

@Component
export default class extends Vue {
  @Prop()
  role: string

  @Prop()
  uid: string

  isError = false
  isLoading = false

  revokeRole() {
    this.isLoading = true
    revokeRole(this.uid, this.role).then(
      () => {
        this.isLoading = false
        this.isError = false
        this.$emit('revoked')
      },
      err => {
        this.isError = true
        this.isLoading = false
        console.error(err)
      },
    )
  }
}
</script>
