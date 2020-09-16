<template>
  <span>
    <input type="date" v-model="dateIso" data-test="input-date" />
    <input type="time" v-model="timeIso" data-test="input-time" />
  </span>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator'
import { dateToISODate, dateToISOTime } from '@/util'

@Component({
  components: {},
})
export default class DateTimeInput extends Vue {
  @Prop()
  value: number

  get dateIso(): string {
    return dateToISODate(new Date(this.value))
  }

  set dateIso(v: string) {
    const newDate = new Date(`${v}T${this.timeIso}`)
    this.$emit('input', newDate.getTime())
  }

  get timeIso(): string {
    return dateToISOTime(new Date(this.value))
  }

  set timeIso(v: string) {
    const newDate = new Date(`${this.dateIso}T${v}`)
    this.$emit('input', newDate.getTime())
  }
}
</script>
