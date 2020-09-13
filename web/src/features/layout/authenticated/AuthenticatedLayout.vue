<template>
  <div>
    <authenticated-header />

    <slot />
  </div>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class AuthenticatedLayout extends Vue {
}
</script>
