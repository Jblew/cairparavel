<template>
  <span>
    <user-role v-for="role of roles" :key="role" :role="role" :uid="uid" />
  </span>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import { projectConfig } from '@/config';
import UserRole from './UserRole.vue'

@Component({
  components: {
    UserRole,
  }
})
export default class extends Vue {
  @Prop()
  uid: string

  get roles(): string[] {
    return projectConfig.roles.list
  }
}
</script>

<style scoped>
</style>
