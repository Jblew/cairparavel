<template>
  <div>
    <authenticated-header />

    <slot />
  </div>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import AuthenticatedHeader from './AuthenticatedHeader.vue'

@Component({
  components: {
    AuthenticatedHeader,
  },
})
export default class AuthenticatedLayout extends Vue {}
</script>
