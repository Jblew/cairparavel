<template>
  <div class="create-event-btn">
    <router-link to="/create-event" data-test="create-event-btn"
      >Create event</router-link
    >
  </div>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'

@Component
export default class extends Vue {}
</script>

<style scoped></style>
