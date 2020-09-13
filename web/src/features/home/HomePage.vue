<template>
  <div>
    <home-header />

    <p>Lorem ipsum</p>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import HomeHeader from './HomeHeader.vue';

@Component({
  components: {
    HomeHeader,
  },
})
export default class HomePage extends Vue {
  @Inject()
  public readonly firebase!: firebase.app.App;

  public user: firebase.User = this.firebase.auth().currentUser!;
}
</script>

<style scoped>
</style>
