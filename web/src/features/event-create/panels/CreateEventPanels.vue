<template>
  <div>
    <state-matches :state="state">
      <template #DetailsSetup
        ><panel-details-setup :machine="machine" :state="state"
      /></template>
      <template #QuestionVoting
        ><panel-question-voting :machine="machine" :state="state"
      /></template>
      <template #VotingSetup
        ><panel-voting-setup :machine="machine" :state="state"
      /></template>
      <template #EventTimeSetup
        ><panel-event-time-setup :machine="machine" :state="state"
      /></template>
      <template #SignupSetup
        ><panel-signup-setup :machine="machine" :state="state"
      /></template>
      <template #ParticipantLimits
        ><panel-participant-limits :machine="machine" :state="state"
      /></template>
      <template #Confirm
        ><panel-confirm :machine="machine" :state="state"
      /></template>
      <template #Success
        ><panel-success :machine="machine" :state="state"
      /></template>
      <template #Error
        ><panel-error :machine="machine" :state="state"
      /></template>
    </state-matches>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator';
import { CreateEventInterpreter } from '../machine'
import { StateMatches } from '@/components'
import PanelConfirm from './PanelConfirm.vue'
import PanelDetailsSetup from './PanelDetailsSetup.vue'
import PanelEventTimeSetup from './PanelEventTimeSetup.vue'
import PanelError from './PanelError.vue'
import PanelParticipantLimits from './PanelParticipantLimits.vue'
import PanelQuestionVoting from './PanelQuestionVoting.vue'
import PanelSignupSetup from './PanelSignupSetup.vue'
import PanelSuccess from './PanelSuccess.vue'
import PanelVotingSetup from './PanelVotingSetup.vue'

@Component({
  components: {
    StateMatches,
    PanelConfirm,
    PanelDetailsSetup,
    PanelError,
    PanelEventTimeSetup,
    PanelParticipantLimits,
    PanelQuestionVoting,
    PanelSignupSetup,
    PanelSuccess,
    PanelVotingSetup,
  }
})
export default class extends Vue {
  @Prop()
  machine: CreateEventInterpreter

  @Prop()
  state: CreateEventInterpreter['state']
}
</script>

<style scoped>
</style>
