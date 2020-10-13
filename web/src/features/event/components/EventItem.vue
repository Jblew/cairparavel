<template>
  <div class="event">
    <state-matches :state="state">
      <template #InitialFetch>
        <event-item-initial-fetch :interpreter="interpreter" :state="state" />
      </template>
      <template #Error>
        <event-item-error :interpreter="interpreter" :state="state" />
      </template>
      <template #TimeVoting>
        <event-item-time-voting :interpreter="interpreter" :state="state" />
      </template>
      <template #DoTimeVote>
        <event-item-do-time-vote :interpreter="interpreter" :state="state" />
      </template>
      <template #DoTimeUnvote>
        <event-item-do-time-vote :interpreter="interpreter" :state="state" />
      </template>
      <template #WaitingForTimeConfirm>
        <event-item-waiting-for-time-confirm
          :interpreter="interpreter"
          :state="state"
        />
      </template>
      <template #DoTimeConfirm>
        <event-item-do-time-confirm :interpreter="interpreter" :state="state" />
      </template>
      <template #Cancelled>
        <event-item-cancelled :interpreter="interpreter" :state="state" />
      </template>
      <template #MembersSignup>
        <event-item-members-signup :interpreter="interpreter" :state="state" />
      </template>
      <template #DoMemberSignup>
        <event-item-do-member-signup
          :interpreter="interpreter"
          :state="state"
        />
      </template>
      <template #DoMemberSignout>
        <event-item-do-member-signout
          :interpreter="interpreter"
          :state="state"
        />
      </template>
      <template #SignupClosed>
        <event-item-signup-closed :interpreter="interpreter" :state="state" />
      </template>
      <template #InProggress>
        <event-item-in-proggress :interpreter="interpreter" :state="state" />
      </template>
      <template #Finished>
        <event-item-finished :interpreter="interpreter" :state="state" />
      </template>
      <template #DoUpdateDetails>
        <event-item-do-update-details
          :interpreter="interpreter"
          :state="state"
        />
      </template>
      <template #DoDelete>
        <event-item-do-delete :interpreter="interpreter" :state="state" />
      </template>
      <template #Deleted>
        <event-item-deleted :interpreter="interpreter" :state="state" />
      </template>
    </state-matches>
    <event-comments :event="event" />
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { Event, EventMachineInterpreter } from '@/businesslogic'
import { getEventInterpreter } from '../eventInterpreter'
import {
  EventItemInitialFetch,
  EventItemError,
  EventItemTimeVoting,
  EventItemDoTimeVote,
  EventItemDoTimeUnvote,
  EventItemWaitingForTimeConfirm,
  EventItemDoTimeConfirm,
  EventItemCancelled,
  EventItemMembersSignup,
  EventItemDoMemberSignup,
  EventItemDoMemberSignout,
  EventItemSignupClosed,
  EventItemInProggress,
  EventItemFinished,
  EventItemDoUpdateDetails,
  EventItemDoDelete,
  EventItemDeleted,
} from './states'
import { StateMatches } from '@/components'
import { EventComments } from '@/features/comments'

@Component({
  components: {
    StateMatches,
    EventItemInitialFetch,
    EventItemError,
    EventItemTimeVoting,
    EventItemDoTimeVote,
    EventItemDoTimeUnvote,
    EventItemWaitingForTimeConfirm,
    EventItemDoTimeConfirm,
    EventItemCancelled,
    EventItemMembersSignup,
    EventItemDoMemberSignup,
    EventItemDoMemberSignout,
    EventItemSignupClosed,
    EventItemInProggress,
    EventItemFinished,
    EventItemDoUpdateDetails,
    EventItemDoDelete,
    EventItemDeleted,
    EventComments,
  },
})
export default class extends Vue {
  @Prop({ required: true })
  eventInitial!: Event

  interpreter: EventMachineInterpreter = getEventInterpreter({ eventId: this.eventInitial.id!, event: this.eventInitial })
  state: EventMachineInterpreter['state'] = this.interpreter.initialState

  created() {
    this.startEventMachine()
  }

  startEventMachine() {
    if (!this.eventInitial.id) throw new Error('Event does not have an ID assigned')
    this.interpreter
      .onTransition(state => {
        this.state = state
      })
      .onEvent(evt => {
        this.onMachineEvent(evt)
      })
      .start()
  }

  beforeDestroy() {
    this.interpreter.stop()
  }

  onMachineEvent(evt: any) {
    if (evt.type === 'ERROR') {
      this.onMachineError(evt)
    }
  }

  onMachineError(errorEvt: { type: 'ERROR', message: string }) {
    alert(`Error: ${errorEvt.message}`)
  }

  get event(): Event {
    return this.state.context.event!
  }
}
</script>

<style scoped></style>
