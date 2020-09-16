<template>
  <div class="authenticated-header">
    <img src="/logo.svg" class="logo" />

    <h1>Hi, {{ user.displayName }}</h1>
    <div class="logout-link-container">
      <logout-link />
    </div>

    <menu-bar />
    <br />
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import LogoutLink from './LogoutLink.vue'
import MenuBar from './MenuBar.vue'

@Component({
  components: {
    LogoutLink,
    MenuBar,
  },
})
export default class AuthenticatedHeader extends Vue {
  @Inject()
  public readonly firebase!: firebase.app.App

  public user: firebase.User = this.firebase.auth().currentUser!
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

.logo {
  height: 3rem;
  float: right;
}

.logout-link-container {
  width: 100%;
}
</style>
