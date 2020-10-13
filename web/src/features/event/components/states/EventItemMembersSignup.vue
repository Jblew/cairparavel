<template>
  <span>
    <event-header :interpreter="interpreter" :state="state" :event="event" />
    <event-path>
      <event-path-item-created :event="event" />
      <event-path-separator />
      <event-path-item-members-signup
        :event="event"
        :signup-enabled="signupEnabled"
        :signout-enabled="signoutEnabled"
        @signup="signup()"
        @signout="signout()"
      />
      <event-path-separator />
      <event-path-item :enabled="true" :checked="false" name="Meeting">
        TODO meeting time
      </event-path-item>
    </event-path>
  </span>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { Event, EventMachineInterpreter } from '@/businesslogic'
import {
  EventHeader,
  EventPath,
  EventPathItem,
  EventPathSeparator,
  EventPathItemCreated,
  EventPathItemMembersSignup,
} from '../components'

@Component({
  components: {
    EventHeader,
    EventPath,
    EventPathItem,
    EventPathSeparator,
    EventPathItemCreated,
    EventPathItemMembersSignup,
  }
})
export default class extends Vue {
  @Prop({ required: true })
  interpreter: EventMachineInterpreter

  @Prop({ required: true })
  state: EventMachineInterpreter['state']

  get event(): Event {
    return this.state.context.event!
  }

  get signupEnabled(): boolean {
    return this.state.nextEvents.includes('SIGNUP_MEMBER')
  }

  signup() {
    this.interpreter.send('SIGNUP_MEMBER')
  }

  get signoutEnabled(): boolean {
    return this.state.nextEvents.includes('SIGNOUT_MEMBER')
  }

  signout() {
    this.interpreter.send('SIGNOUT_MEMBER')
  }
}
</script>

<style scoped></style>
