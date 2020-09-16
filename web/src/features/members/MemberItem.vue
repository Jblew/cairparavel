<template>
  <span>
    <strong>{{ email }}</strong>
    Roles: <user-roles :uid="uid" />
  </span>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { UserDoc } from './fetchMembers'
import { UserRoles } from './roles'

@Component({
  components: {
    UserRoles,
  },
})
export default class extends Vue {
  @Prop()
  member: UserDoc

  get email(): string {
    return this.member.Email
  }

  get uid(): string {
    return this.member.UID
  }
}
</script>

<style scoped>
span {
  text-align: left;
}
</style>
