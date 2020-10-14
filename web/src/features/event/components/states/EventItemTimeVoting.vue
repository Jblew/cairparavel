<template>
  <span>
    <event-header :interpreter="interpreter" :state="state" :event="event" />
    <event-path>
      <event-path-item-created :event="event" />
      <event-path-separator />
      <event-path-item :enabled="true" :checked="true" name="Voting for time">
        <date-voting
          :can-add="canAdd"
          :available="available"
          :selected="myVotes"
          :votes="votesDisplayNames"
          @vote="vote"
          @unvote="unvote"
        />
      </event-path-item>
      <event-path-separator />
      <event-path-item :enabled="false" :checked="false" name="Signup" />
      <event-path-separator />
      <event-path-item-meeting-unknown-time
        :enabled="false"
        :checked="false"
        :event="event"
      />
    </event-path>
  </span>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { Event, EventMachineInterpreter } from '@/businesslogic'
import { DateVoting } from '@/components'
import {
  EventHeader,
  EventPath,
  EventPathItem,
  EventPathSeparator,
  EventPathItemCreated,
  EventPathItemMeetingUnknownTime,
} from '../components'

@Component({
  components: {
    EventHeader,
    EventPath,
    EventPathItem,
    EventPathSeparator,
    EventPathItemCreated,
    EventPathItemMeetingUnknownTime,
    DateVoting,
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

  get canAdd(): boolean {
    return this.event.canSuggestTime
  }

  get available(): number[] {
    return this.event.allowedTimes
  }

  get myVotes(): number[] {
    const uid = this.state.context.currentUid
    const myVotes = this.event.votes?.[uid]
    if (!myVotes) return []
    return myVotes.times
  }

  get votesDisplayNames(): Record<number, string[]> {
    const votesByUid = this.event.votes || {}
    const getDisplayNamesForTime = (time: number): string[] => {
      return Object.values(votesByUid).filter(vote => vote.times.includes(time)).map(vote => vote.displayName)
    }

    return this.available.reduce((byTime, time) => {
      return {
        ...byTime,
        [time]: getDisplayNamesForTime(time),
      }
    }, {} as Record<number, string[]>)
  }

  vote(time: number) {
    console.log('vote', time)
    this.interpreter.send({ type: 'TIME_VOTE', time })
  }

  unvote(time: number) {
    this.interpreter.send({ type: 'TIME_UNVOTE', time })
  }
}
</script>

<style scoped></style>
