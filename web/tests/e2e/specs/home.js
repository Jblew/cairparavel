import { addSampleEvent } from '../helpers/sample-event.js'

describe('Home page', () => {
  beforeEach(() => cy.login() && cy.setCookie('e2e', '1'))

  it('Shows home page', () => {
    cy.visit('/')
    cy.contains('h1', 'Hi')
  })

  it('Shows menu', () => {
    cy.visit('/')
    cy.contains('.router-link-exact-active', 'Home')
  })

  it('Shows create event button', () => {
    cy.visit('/')
    cy.contains('[data-test="create-event-btn"]', 'Create event')
  })

  it('Shows events', () => {
    const name = 'Event #' + Date.now()
    addSampleEvent({ name })
    cy.visit('/')
    cy.get('[data-test="events-list"]').should('be.visible')
    cy.get('[data-test="events-list"]').find('.event').its('length').should('be.gte', 1)
  })

  it('Newest event is at the top of the list', () => {
    const name = 'Event #' + Date.now()
    addSampleEvent({ name })
    cy.visit('/')
    cy.get('[data-test="events-list"]').should('be.visible')
    cy.get('[data-test="events-list"] .event').first().get('.name').contains(name)
  })
})
