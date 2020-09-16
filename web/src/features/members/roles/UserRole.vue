<template>
  <a v-if="isError">Error checking role {{ role }}</a>
  <a v-else-if="isLoading">Checking role {{ role }}</a>
  <revoke-role-button
    v-else-if="hasRole"
    :role="role"
    :uid="uid"
    @revoked="recheckRole"
  />
  <grant-role-button v-else :role="role" :uid="uid" @granted="recheckRole" />
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { projectConfig } from '@/config'
import { checkHasRole } from './helpers'
import GrantRoleButton from './GrantRoleButton.vue'
import RevokeRoleButton from './RevokeRoleButton.vue'

@Component({
  components: {
    RevokeRoleButton,
    GrantRoleButton,
  },
})
export default class extends Vue {
  @Prop()
  role: string

  @Prop()
  uid: string

  hasRole = false
  isError = false
  isLoading = false

  mounted() {
    this.recheckRole()
  }

  recheckRole() {
    this.isLoading = true
    checkHasRole(this.uid, this.role).then(
      hasRole => {
        this.hasRole = hasRole
        this.isLoading = false
        this.isError = false
      },
      err => {
        this.isError = true
        this.hasRole = false
        this.isLoading = false
        console.error(err)
      },
    )
  }
}
</script>

<style scoped>
a {
  border: 1px solid #ccc;
  padding: 0.25rem 0.5rem 0.25rem 0.5rem;
  margin-right: 0.5rem;
  display: inline-block;
  margin-bottom: 0.5rem;
}

a.clickable {
  text-decoration: underline;
  cursor: pointer;
}
</style>
