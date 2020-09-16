describe('Make sure cypress firebase integration works', () => {
  it('Adds document to test_hello_world collection of Firestore', () => {
    cy.callFirestore('add', 'test/cypress-integration/' + Date.now(), {
      some: 'value',
    })
  })
})
