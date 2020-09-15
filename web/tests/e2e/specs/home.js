describe("Home page", () => {
  beforeEach(() => cy.login())

  it("Shows home page", () => {
    cy.visit("/");
    cy.contains("h1", "Hi");
  })

  it("Shows menu", () => {
    cy.visit("/");
    cy.contains(".router-link-exact-active", "Home");
  })

  it("Shows events", () => {
    cy.visit("/");
    cy.contains(".router-link-exact-active", "Home");
  })

  it("Shows create event button", () => {
    cy.visit("/");
    cy.contains('[data-test="create-event-btn"]', "Create event");
  })
})
