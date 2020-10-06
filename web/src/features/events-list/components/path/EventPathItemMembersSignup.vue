<template>
  <event-path-item :enabled="true" :checked="true" name="Signup">
    <ul>
      <li v-for="displayName in signedMembersDisplayNames" :key="displayName">
        {{ displayName }}
      </li>
    </ul>
    <button v-if="signupEnabled" @click="$emit('signup')">Sign up</button>
    <button v-if="signoutEnabled" @click="$emit('signout')">Sign out</button>
    <span>
      Signed in {{ currentParticipants }} of max
      {{ maxParticipants }} participants. (Minimum {{ minParticipants }} to
      start event)
    </span>
  </event-path-item>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { Event, EventSignup } from '@/businesslogic'
import EventPathItem from './EventPathItem.vue'

@Component({
  components: {
    EventPathItem,
  }
})
export default class extends Vue {
  @Prop({ required: true })
  event: Event

  @Prop({ required: true })
  signupEnabled: boolean

  @Prop({ required: true })
  signoutEnabled: boolean

  get signedMembers(): Record<string, EventSignup> {
    return this.event.signedMembers || {}
  }

  get signedMembersDisplayNames(): string[] {
    return Object.values(this.signedMembers).map(signup => signup.displayName)
  }

  get minParticipants(): number {
    return this.event.minParticipants
  }

  get maxParticipants(): number {
    return this.event.maxParticipants
  }

  get currentParticipants(): number {
    return Object.keys(this.signedMembers).length
  }
}
</script>

<style scoped></style>
