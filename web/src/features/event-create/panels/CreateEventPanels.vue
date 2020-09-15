<template>
  <div>
    <state-matches :state="state">
      <template #DetailsSetup><panel-details-setup /></template>
      <template #QuestionVoting><panel-question-voting /></template>
      <template #VotingSetup><panel-voting-setup /></template>
      <template #EventTimeSetup><panel-event-time-setup /></template>
      <template #SignupSetup><panel-signup-setup /></template>
      <template #ParticipantLimits><panel-participant-limits /></template>
      <template #Confirm><panel-confirm /></template>
      <template #Success><panel-success /></template>
      <template #Error><panel-error /></template>
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
  @Inject()
  machine: CreateEventInterpreter

  @Inject()
  state: CreateEventInterpreter['state']
}
</script>

<style scoped>
</style>
