<template>
  <div class="logout-link">
    <router-link to="/logout">Log me out</router-link>
  </div>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'

@Component
export default class extends Vue {}
</script>

<style scoped>
.logout-link {
  text-align: left;
  font-size: 75%;
  float: right;
}

.logout-link a {
  display: block;
}
</style>
