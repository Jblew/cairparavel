<template>
  <div class="authenticated-header">
    <img src="/logo.svg" class="logo" />

    <h1>Hi, {{ user.displayName }}</h1>
    <div class="logout-link">
      <router-link to="/logout">Log me out</router-link>
    </div>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';

@Component
export default class AuthenticatedHeader extends Vue {
  @Inject()
  public readonly firebase!: firebase.app.App;

  public user: firebase.User = this.firebase.auth().currentUser!;
}
</script>

<style scoped>
.authenticated-header {
  text-align: left;
  padding: 1rem;
  background: #eee;
  margin-bottom: 2rem;
}

.authenticated-header h1 {
  margin: 0;
}

.logout-link {
  text-align: left;
  width: 100%;
  font-size: 75%;
}

.logo {
  height: 3rem;
  float: right;
}
</style>
