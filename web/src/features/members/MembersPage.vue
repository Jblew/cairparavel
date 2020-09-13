<template>
  <authenticated-layout>
    <p>Members</p>
  </authenticated-layout>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import { AuthenticatedLayout } from '@/features/layout';

@Component({
  components: {
    AuthenticatedLayout,
  },
})
export default class MembersPage extends Vue {

}
</script>

<style scoped>
</style>
