describe('Plant Profiles', () => {
    beforeEach(() => {
        // Cypress starts out with a blank slate for each test,
        // so we must tell it to visit our website with the `cy.visit()` command.
        // Since we want to visit the same URL at the start of all our tests,
        // we include it in our beforeEach function so that it runs before each test
        cy.visit('http://localhost:8090')
    })

    it('adds and removes a plant profile', () => {
        cy.get('h1 a').should('have.text', 'Plant List')

        cy.get('nav ul li')
            .eq(1)
            .should('have.text', 'Add a Plant')
            .click()

        cy.get('input[type=text]')
            .eq(0)
            .should('have.attr', 'name', 'name')
            .type("Latin Test Plant Name")

        cy.get('input[type=text]')
            .eq(1)
            .should('have.attr', 'name', 'common_name')
            .type("Test Plant")

        cy.get('input[type=text]')
            .eq(2)
            .should('have.attr', 'name', 'seed_company')
            .type("a Test Seed Company")

        cy.get('input[type=number]')
            .eq(0)
            .should('have.attr', 'name', 'expected_days_to_harvest')
            .type("50")

        cy.get('[id=type]')
            .eq(0)
            .should('have.attr', 'name', 'type')
            .select("Harvest Once")

        cy.get('input[type=number]')
            .eq(1)
            .should('have.attr', 'name', 'ph_low')
            .type("1.2")

        cy.get('input[type=number]')
            .eq(2)
            .should('have.attr', 'name', 'ph_high')
            .type("1.3")

        cy.get('input[type=number]')
            .eq(3)
            .should('have.attr', 'name', 'ec_low')
            .type("1.4")

        cy.get('input[type=number]')
            .eq(4)
            .should('have.attr', 'name', 'ec_high')
            .type("1.5")

        cy.get('button')
            .should('have.text', 'Submit')
            .click()

        // check the list of plants to see if the test plant is there
        cy.get("table")
            .find("tr")
            .then((row) => {
                //row.length will give you the row count
                cy.log(row.length);
            });

        const row2 = cy.get("table")
            .find("tr")
            .eq(1)

        let tdIndex = 1
        row2.get('td').eq(tdIndex++)
            .should('have.text', 'Latin Test Plant Name')

        row2.get("td").eq(tdIndex++)
            .should('have.text', 'Test Plant')

        row2.get("td").eq(tdIndex++)
            .should('have.text', 'a Test Seed Company')

        row2.get('td').eq(tdIndex++)
            .should('have.text', '50')

        row2.get('td').eq(tdIndex++)
            .should('have.text', 'harvest_once')

        // go delete the test plant
        cy.get('nav ul li')
            .eq(2)
            .should('have.text', 'Delete a Plant')
            .click()

        cy.get('input[type=text]')
            .eq(0)
            .should('have.attr', 'name', 'id')
            .type("1")

        cy.get('button')
            .should('have.text', 'Delete')
            .click()
    })
})