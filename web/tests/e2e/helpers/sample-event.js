const { DateTime } = require('luxon');
const { v4: uuid } = require('uuid');

function getSampleEvent(modifiers) {
  const startTime = DateTime.local().plus({ years: 15 })
  const endTime = startTime.plus({ hours: 1 })
  const signupTime = startTime.minus({ days: 1 })
  const ownerUid = process.env.CYPRESS_TEST_UID
  return {
    startTime: startTime.toMillis(),
    endTime: endTime.toMillis(),
    signupTime: signupTime.toMillis(),
    votingTime: - 1,
    allowedTimes: [],
    canSuggestTime: true,
    name: 'Test event with a name',
    description: 'Test event with a description',
    minParticipants: 1,
    maxParticipants: 5,
    ownerDisplayName: 'Test event owner',
    ownerUid,
    timeConfirmed: true,
    ...modifiers
  }
}

function addSampleEvent(modifiers) {
  const id = uuid();
  const data = getSampleEvent(modifiers)
  cy.callFirestore('set', 'envs/test/events/' + id, data)
  return id
}


module.exports = { getSampleEvent, addSampleEvent }
