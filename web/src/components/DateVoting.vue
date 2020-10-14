<template>
  <div class="date-voting">
    {{ selected }}
    <b-form-group label="Vote">
      <b-form-checkbox
        v-for="time in available"
        :key="time"
        :checked="isSelected(time)"
        @change="toggle(time)"
      >
        <span v-b-tooltip.hover :title="participantsTooltipContents(time)">
          {{ time | formatTimeMillis }} +{{ numberOfVotes(time) }}
        </span>
      </b-form-checkbox>
    </b-form-group>
    <span v-if="canAdd">
      <date-time-input v-model="addDateValue" />
      <b-button @click="addDate">Add</b-button>
    </span>
  </div>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { DateTimeInput } from './inputs'

function formatTimeMillis(timeMs: number): string {
  return new Date(timeMs).toISOString()
}

@Component({
  filters: {
    formatTimeMillis,
  },
  components: {
    DateTimeInput,
  },
})
export default class extends Vue {
  @Prop({ required: true })
  votes: Record<number, string[]>

  @Prop({ required: true })
  available: number[]

  @Prop({ required: true })
  selected: number[]

  @Prop({ required: true })
  canAdd: boolean

  addDateValue = Date.now()

  isSelected(time: number): boolean {
    return this.selected.includes(time)
  }

  toggle(time: number) {
    const isSelected = this.selected.includes(time)
    if (isSelected) {
      this.$emit('unvote', time)
    } else {
      this.$emit('vote', time)
    }
  }

  numberOfVotes(time: number): number {
    return this.votes[time].length
  }

  participantsTooltipContents(time: number): string {
    return this.votes[time].join("\n")
  }

  addDate() {
    this.$emit('vote', this.addDateValue)
  }
}

interface Model {

}
</script>
<style>
</style>
