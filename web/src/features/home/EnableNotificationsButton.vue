<template>
  <div class="create-event-btn">
    <a :href="messengerLink" target="_blank">
      Enable messenger notifications
    </a>
  </div>
</template>
<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { projectConfig } from '../../config'
import firebase from 'firebase/app'

@Component
export default class extends Vue {
  get messengerLink(): string {
    return projectConfig.messengerChatUrl(firebase.auth().currentUser?.uid || '')
  }
}
</script>

<style scoped></style>
