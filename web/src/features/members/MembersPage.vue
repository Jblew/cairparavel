<template>
  <authenticated-layout>
    <stateful-resource :resource="membersResource">
      <ul>
        <li v-for="member in members" :key="member.uid">
          <member-item :member="member" />
        </li>
      </ul>
    </stateful-resource>
  </authenticated-layout>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { AuthenticatedLayout } from '@/features/layout'
import { StatefulResource, Resource } from 'vue-stateful-resource'
import { UserDoc, fetchMembers } from './fetchMembers'
import MemberItem from './MemberItem.vue'

@Component({
  components: {
    AuthenticatedLayout,
    StatefulResource,
    MemberItem,
  },
})
export default class MembersPage extends Vue {
  membersResource: Resource<UserDoc[]> = Resource.empty()

  mounted() {
    Resource.fetchResource(
      'members',
      fetchMembers,
      res => (this.membersResource = res),
    )
  }

  get members(): UserDoc[] {
    return this.membersResource.result || []
  }
}
</script>

<style scoped>
ul li {
  text-align: left;
}
</style>
