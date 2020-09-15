describe("Create event", () => {
  beforeEach(() => cy.login())

  it("Creates event with voting", () => {
    const eventName = "Test event " + Date.now()
    cy.visit("/#/create-event");

    cy.get('[data-test="panel-details"]').should('be.visible')
    cy.get('[data-test="input-name"]').type(eventName)
    cy.get('[data-test="input-description"]').type("This is a test event")
    cy.get('[data-test="btn-next"]').click()

    cy.get('[data-test="panel-voting-question"]').should('be.visible')
    cy.get('[data-test="btn-choose-voting"]').click()

    cy.get('[data-test="panel-voting-setup"]').should('be.visible')
    cy.get('[data-test="input-can-suggest-time"]').click()
    cy.get('[data-test="input-allowed-time-0"]').type('2030-05-05 17:00')
    cy.get('[data-test="btn-add-allowed-time"]').click()
    cy.get('[data-test="input-allowed-time-1"]').type('2030-05-04 1:00')
    cy.get('[data-test="btn-add-allowed-time"]').click()
    cy.get('[data-test="input-allowed-time-2"]').type('2030-05-06 17:00')
    cy.get('[data-test="input-voting-time"]').type('2030-04-15 00:00')
    cy.get('[data-test="btn-next"]').click()

    cy.get('[data-test="panel-signup-time"]').should('be.visible')
    cy.get('[data-test="input-signup-time"]').type('2030-04-20 17:00')
    cy.get('[data-test="btn-next"]').click()

    cy.get('[data-test="panel-participants-limits"]').should('be.visible')
    cy.get('[data-test="btn-back"]').click()
    cy.get('[data-test="btn-next"]').click()

    cy.get('[data-test="panel-participants-limits"]').should('be.visible')
    cy.get('[data-test="input-min-participants"]').type('5')
    cy.get('[data-test="input-max-participants"]').type('10')
    cy.get('[data-test="btn-next"]').click()

    cy.get('[data-test="panel-confirm"]').should('be.visible')
    cy.get('[data-test="btn-next"]').click()

    cy.get('[data-test="panel-success"]').should('be.visible')
  })
})
