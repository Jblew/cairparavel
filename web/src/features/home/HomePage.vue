<template>
  <authenticated-layout>
    <enable-notifications-button />
    <create-event-button />
    <h2>Events</h2>
    <events-list />
  </authenticated-layout>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { AuthenticatedLayout } from '@/features/layout'
import { EventsList } from '@/features/events-list'
import CreateEventButton from './CreateEventButton.vue'
import EnableNotificationsButton from './EnableNotificationsButton.vue'

@Component({
  components: {
    AuthenticatedLayout,
    CreateEventButton,
    EnableNotificationsButton,
    EventsList,
  },
})
export default class HomePage extends Vue {
  @Inject()
  public readonly firebase!: firebase.app.App

  public user: firebase.User = this.firebase.auth().currentUser!
}
</script>

<style scoped></style>
