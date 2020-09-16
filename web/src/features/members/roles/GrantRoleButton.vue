<template>
  <a v-if="isLoading">Granting {{ role }}...</a>
  <a v-else-if="isError" @click="grantRole" class="clickable"
    >Grant {{ role }} (error!)</a
  >
  <a v-else @click="grantRole" class="clickable">Grant {{ role }}</a>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { grantRole } from './helpers'

@Component
export default class extends Vue {
  @Prop()
  role: string

  @Prop()
  uid: string

  isError = false
  isLoading = false

  grantRole() {
    this.isLoading = true
    grantRole(this.uid, this.role).then(
      () => {
        this.isLoading = false
        this.isError = false
        this.$emit('granted')
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
