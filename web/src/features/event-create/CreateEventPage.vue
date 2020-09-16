<template>
  <authenticated-layout>
    <create-event-form />
  </authenticated-layout>
</template>
<script lang="ts">
import { Component, Prop, Vue, Inject } from 'vue-property-decorator'
import { AuthenticatedLayout } from '@/features/layout'
import CreateEventForm from './CreateEventForm.vue'

@Component({
  components: {
    AuthenticatedLayout,
    CreateEventForm,
  },
})
export default class CreateEventPage extends Vue {}
</script>

<style scoped></style>
